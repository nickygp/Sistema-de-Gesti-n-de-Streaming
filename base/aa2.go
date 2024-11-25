/*
@autor: Nicolas Garcia
@fecha 21/11/2024
@descripcion: Aprendizaje Autonomo 2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Definir categorías
	categories := []string{
		"Películas",
		"Series",
		"Documentales",
		"Infantiles",
		"Salir",
	}

	for {
		// Mostrar el menú
		fmt.Println("¡Bienvenido, a continuación te presentaremos las categorías " +
			"que podrás disfrutar!\n\n--- Menú de Categorías ---")
		for i, category := range categories {
			fmt.Printf("%d. %s\n", i+1, category)
		}

		// Elección del usuario
		fmt.Print("\nSelecciona un número de acuerdo a tu preferencia: ")
		reader := bufio.NewReader(os.Stdin)
		var choice int
		_, err := fmt.Fscanf(reader, "%d\n", &choice)
		if err != nil || choice < 1 || choice > len(categories) {
			fmt.Println("Por favor, elige una opción válida.")
			reader.ReadString('\n') // Limpia la entrada si hay errores
			continue
		}

		// Ver opción seleccionada
		if choice == len(categories) {
			fmt.Println("Te dirigirás a la pantalla de inicio...")
			break
		} else {
			fmt.Printf("Te dirigiremos hacia el contenido de: %s\n", categories[choice-1])
		}
	}
}
