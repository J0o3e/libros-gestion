package main

import (
	"fmt"

	"github.com/joesolano/libros-gestion/services"
)

func main() {
	for {
		fmt.Println("============ Sistema de Gestion de Libros Electronicos =======")
		fmt.Println("1. Gestionar libros")
		fmt.Println("2. Gestionar usuarios")
		fmt.Println("3. Gestionar prestamos")
		fmt.Println("4. Salir")
		fmt.Println("Seleccione una opcion: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			menuLibros()
		case 2:
			menuUsarios()
		case 3:
			menuPrestamos()
		case 4:
			fmt.Println("Gracias por usar el sistema!")
			return
		default:
			fmt.Println("Opcion invalida, intente de nuevo")
		}
	}
}

func menuLibros() {                                         //Funcion de libros//
	for {
		fmt.Println("\n-- Gestion de Libros---")
		fmt.Println("1. Añadir libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Volver al menu principal")
		fmt.Println("Seleccione una opcion: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			services.AddBook()
		case 2:
			services.ListBooks()
		case 3:
			return
		default:
			fmt.Println("Opcion Invalida")
		}

	}
}

func menuUsarios() {                                          //Funcion de Usuarios//
	for {
		fmt.Println("\n--- Gestion De Usuarios ---")
		fmt.Println("1. Registrar usuario")
		fmt. Println("2. Listar usuarios")
		fmt.Println("3. Volver al menu principal")
		fmt.Println("Seleccione una opcion: ")

		var opcion int 
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			services.RegisterUser()
		case 2:
			services.ListUsers()
		case 3:
			return
		default:
			fmt.Println("Opcion invalida.")
		}
	} 
}

func menuPrestamos() {                                     //Funcion de Prestamos//
	for {
		fmt.Println("\n-- Gestion de Prestamos ---")
		fmt.Println("1. Prestar Libro")
		fmt.Println("3. Listar prestamos")
		fmt.Println("4, Volver al menu principal")
		fmt.Println("Seleccione una opcion: ")
		
		var opcion int
		fmt.Scanln(&opcion)
		
		switch opcion {
		case 1:
			services.PrestarLibro()
		case 2:
			services.DevolverLibro()
		case 3: 
			services.ListarPrestamos()
		case 4: 
			return
		default:
			fmt.Println("Opcion Invalida.")
		}
	}
}
