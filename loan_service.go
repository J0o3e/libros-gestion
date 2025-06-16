package services

import (
	"fmt"
	"time"

	"github.com/joesolano/libros-gestion/models"
)

var Prestamos []models.Loan

func PrestarLibro() {
	var libroID, usuarioID string
	fmt.Print("ID del libro: ")
	fmt.Scanln(&libroID)

	fmt.Print ("ID del usuario: ")
	fmt.Scanln(&usuarioID)

	                                 //Verificar si el libro esta disponible
	for i, libro := range Libros{
		if libro.ID == libroID {
			if libro.Disponible {
				                     // Marcar como no disponible
				Libros[i].Disponible = false

									// Creacion de prestamos
				id := fmt.Sprintf("P%d", time.Now().UnixNano())
				prestamo := models.Loan{
					ID: id,
					LibroID: libroID,
					UsuarioID: usuarioID,
					Inicio: time.Now(),
					Fin: nil,
					Estado: "prestado",
				}

				Prestamos = append(Prestamos, prestamo)
				fmt.Println("Libro prestado exitosamente.")
				return
			} else {
				fmt.Println("El libro no esta disponible.")
				return
			}
		}
	}

	fmt.Println("Libro no encontrado")
}

func DevolverLibro() {
	var prestamoID string
	fmt.Print("ID del prestamo: ")
	fmt.Scanln(&prestamoID)

	for i, p := range Prestamos {
		if p.ID == prestamoID && p.Estado == "prestado" {
			now := time.Now()
			Prestamos[i].Estado = "devuelto"
			Prestamos[i].Fin = &now

			                                                //Marcar libro como disponible otra vez
			for j := range Libros {
				if Libros[j].ID == p.LibroID {
					Libros[j].Disponible = true
					break
				}
			}

			fmt.Println("Libro devuelto con exito.")
			return
		}
	}

	fmt.Println("Prestamo no encontrado o ya devuelto")
}

func ListarPrestamos() {
	if len(Prestamos) == 0 {
		fmt.Println("No hay prestamos registrados.")
		return
	}

	fmt.Println("Historial de prestamos: ")
	for _, p := range Prestamos {
		fin := "En curso"
		if p.Fin != nil {
			fin = p.Fin.Format("2006-01-02 15:04")
		}

		fmt.Printf("- %s | libro: %s | Usuario: %s | Inicio: %s | Fin: %s | Estado: %s\n",
			p.ID, p.LibroID, p.UsuarioID, p.Inicio.Format("2006-01-02 15:04"), fin, p.Estado)
	}
}
