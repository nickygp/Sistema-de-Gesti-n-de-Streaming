/*
@autor: Nicolas Garcia
@fecha 02/11/2024
@descripcion: Aprendizaje Autonomo 3
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Sistema es la estructura principal que gestiona el menú y las categorías
type Sistema struct {
	categorias map[string][]string
}

// El NewSistema crea una nueva instancia de la estructura Sistema
func NewSistema() *Sistema {
	return &Sistema{
		categorias: make(map[string][]string),
	}
}

// SetCategoria agrega o actualiza una categoría y sus opciones
func (s *Sistema) SetCategoria(nombre string, opciones []string) {
	s.categorias[nombre] = opciones
}

// GetCategorias devuelve la lista de categorías
func (s *Sistema) GetCategorias() []string {
	categorias := make([]string, 0, len(s.categorias))
	for categoria := range s.categorias {
		categorias = append(categorias, categoria)
	}
	return categorias
}

// GetOpciones devuelve las opciones de una categoría seleccionada
func (s *Sistema) GetOpciones(categoria string) ([]string, bool) {
	opciones, existe := s.categorias[categoria]
	return opciones, existe
}

// MostrarMenúPrincipal muestra el menú principal y realiza la interacción con el usuario
func (s *Sistema) MostrarMenúPrincipal() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Se muestran las categorías
		fmt.Println("¡Bienvenido, a continuación te presentaremos las categorías " +
			"que podrás disfutar!\n\n--- Menú de Categorías ---")
		categorias := s.GetCategorias()
		for i, categoria := range categorias {
			fmt.Printf("%d. %s\n", i+1, categoria)
		}
		fmt.Printf("%d. Salir\n", len(categorias)+1)

		// Leer la elección del usuario
		fmt.Print("\nSelecciona un número de acuerdo a tu preferencia: ")
		var choice int
		_, err := fmt.Fscanf(reader, "%d\n", &choice)
		if err != nil || choice < 1 || choice > len(categorias)+1 {
			fmt.Println("Por favor, elige una opción válida.")
			reader.ReadString('\n') // Limpia la entrada si hay errores
			continue
		}

		// Opción 5 de salir
		if choice == len(categorias)+1 {
			fmt.Println("¡Adiós, nos vemos pronto!")
			break
		}

		// Mostrar todas las opciones dentro de la categoría seleccionada
		selectedCategory := categorias[choice-1]
		fmt.Printf("\n--- Opciones en %s ---\n", selectedCategory)
		opciones, _ := s.GetOpciones(selectedCategory)
		for i, opcion := range opciones {
			fmt.Printf("%d. %s\n", i+1, opcion)
		}
		fmt.Println("\nPresiona Enter para regresar al menú principal...")
		reader.ReadString('\n')
	}
}

func main() {
	// Función main que crea el sistema y agregar opciones de las categorías
	sistema := NewSistema()
	sistema.SetCategoria("Películas", []string{"Acción", "Drama", "Terror", "Comedia"})
	sistema.SetCategoria("Series", []string{"Ciencia Ficción", "Juveniles", "Acción", "Intriga"})
	sistema.SetCategoria("Documentales", []string{"Deporte", "Historia", "Alimentación", "Empresas"})
	sistema.SetCategoria("Infantiles", []string{"Animación", "Educativos", "En Familia", "Didáctico"})

	// Mostrar el menú principal
	sistema.MostrarMenúPrincipal()
}
