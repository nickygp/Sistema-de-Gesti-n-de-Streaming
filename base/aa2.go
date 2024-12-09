/*
@autor: Nicolas Garcia
@fecha 02/12/2024
@descripción: Aprendizaje Autónomo 3
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Estructura Usuario
type Usuario struct {
	id     int
	nombre string
	correo string
}

// Getters y Setters para Usuario
func (u *Usuario) GetID() int {
	return u.id
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

// Error introducir usuario
func (u *Usuario) SetNombre(nombre string) error {
	if len(strings.TrimSpace(nombre)) == 0 {
		return errors.New("El campo nombre no puede estar vacío.")
	}
	u.nombre = nombre
	return nil
}

func (u *Usuario) GetCorreo() string {
	return u.correo
}

// Erro introducir correo
func (u *Usuario) SetCorreo(correo string) error {
	if !strings.Contains(correo, "@") || len(strings.TrimSpace(correo)) < 5 {
		return errors.New("Correo inválido. (Ejemplo válido: usuario@ejemplo.com)")
	}
	u.correo = correo
	return nil
}

// Estructura Contenido
type Contenido struct {
	Id        int
	titulo    string
	categoria string
}

// Getters y Setters para Contenido
func (c *Contenido) GetID() int {
	return c.Id
}

func (c *Contenido) GetTitulo() string {
	return c.titulo
}

func (c *Contenido) SetTitulo(titulo string) error {
	if len(strings.TrimSpace(titulo)) == 0 {
		return errors.New("El campo título no puede estar vacío")
	}
	c.titulo = titulo
	return nil
}

func (c *Contenido) GetCategoria() string {
	return c.categoria
}

func (c *Contenido) SetCategoria(categoria string) error {
	if len(strings.TrimSpace(categoria)) == 0 {
		return errors.New("El campo categoría no puede estar vacío")
	}
	c.categoria = categoria
	return nil
}

// Estructura SistemaStreaming
type SistemaStreaming struct {
	usuarios           []Usuario
	contenidos         []Contenido
	proximoIDUsuario   int
	proximoIDContenido int
}

// Constructor para el sistema
func NuevoSistemaStreaming() *SistemaStreaming {
	return &SistemaStreaming{
		usuarios:           []Usuario{},
		contenidos:         []Contenido{},
		proximoIDUsuario:   1,
		proximoIDContenido: 1,
	}
}

// Métodos para agregar y listar usuarios
func (s *SistemaStreaming) AgregarUsuario(nombre, correo string) error {
	for _, u := range s.usuarios {
		if u.correo == correo {
			return errors.New("Correo registrado.")
		}
	}

	usuario := Usuario{id: s.proximoIDUsuario}
	if err := usuario.SetNombre(nombre); err != nil {
		return err
	}
	if err := usuario.SetCorreo(correo); err != nil {
		return err
	}

	s.usuarios = append(s.usuarios, usuario)
	s.proximoIDUsuario++
	return nil
}

func (s *SistemaStreaming) ListarUsuarios() {
	if len(s.usuarios) == 0 {
		fmt.Println("No existen usuarios registrados.")
		return
	}
	fmt.Println("\n--- Lista de Usuarios ---")
	for _, u := range s.usuarios {
		fmt.Printf("ID: %d | Nombre: %s | Correo: %s\n", u.GetID(), u.GetNombre(), u.GetCorreo())
	}
}

// Métodos para agregar y listar contenidos
func (s *SistemaStreaming) AgregarContenido(titulo, categoria string) error {
	contenido := Contenido{Id: s.proximoIDContenido}
	if err := contenido.SetTitulo(titulo); err != nil {
		return err
	}
	if err := contenido.SetCategoria(categoria); err != nil {
		return err
	}

	s.contenidos = append(s.contenidos, contenido)
	s.proximoIDContenido++
	return nil
}

func (s *SistemaStreaming) ListarContenidos() {
	if len(s.contenidos) == 0 {
		fmt.Println("No existen contenidos disponibles.")
		return
	}
	fmt.Println("\n--- Lista de Contenidos ---")
	for _, c := range s.contenidos {
		fmt.Printf("ID: %d | Título: %s | Categoría: %s\n", c.GetID(), c.GetTitulo(), c.GetCategoria())
	}
}

// Función main
func main() {
	sistema := NuevoSistemaStreaming()
	lector := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("¡Bienvenid@, a continuación te presentaremos las categorías " +
			"que podrás seleccionar!\n\n--- Menú de Categorías ---")
		fmt.Println("1. Añadir Usuario")
		fmt.Println("2. Listar Usuarios")
		fmt.Println("3. Añadir Contenido")
		fmt.Println("4. Listar Contenidos")
		fmt.Println("5. Salir")
		fmt.Print("Selecciona una opción: ")

		entrada, _ := lector.ReadString('\n')
		eleccion := strings.TrimSpace(entrada)

		switch eleccion {
		case "1":
			fmt.Print("Introduce el nombre de usuario: ")
			nombre, _ := lector.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Introduce el correo del usuario: ")
			correo, _ := lector.ReadString('\n')
			correo = strings.TrimSpace(correo)

			if err := sistema.AgregarUsuario(nombre, correo); err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Println("Usuario añadido correctamente.")
			}

		case "2":
			sistema.ListarUsuarios()

		case "3":
			fmt.Print("Introduce el título del contenido: ")
			titulo, _ := lector.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			fmt.Print("Introduce la categoría del contenido: ")
			categoria, _ := lector.ReadString('\n')
			categoria = strings.TrimSpace(categoria)

			if err := sistema.AgregarContenido(titulo, categoria); err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Println("Contenido añadido correctamente.")
			}

		case "4":
			sistema.ListarContenidos()

		case "5":
			fmt.Println("¡Adiós! Nos vemos pronto.")
			return

		default:
			fmt.Println("Opción inválida, inténtalo nuevamente.")
		}
	}
}
