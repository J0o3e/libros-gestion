package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joesolano/libros-gestion/models"
)

var Libros []models.Book

func AddBook() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Titulo: ")
	titulo, _ := reader.ReadString('\n')

	fmt.Print("Autor: ")
	autor, _ := reader.ReadString('\n')

	fmt.Print("Genero: ")
	genero, _ := reader.ReadString('\n')

	fmt.Print("Año de publicacion: ")
	yearStr, _ := reader.ReadString('\n')
	yearStr = strings.TrimSpace(yearStr)
	year, _ := strconv.Atoi(yearStr)

	id := fmt.Sprintf("B%d", time.Now().UnixNano())

	libro := models.Book{
		ID:         id,
		Titulo:     strings.TrimSpace(titulo),
		Autor:      strings.TrimSpace(autor),
		Genero:     strings.TrimSpace(genero),
		Year:       year,
		Disponible: true,
	}

	Libros = append(Libros, libro)
	fmt.Println("Libro añadido exitosamente")

}

func ListBooks() {
	if len(Libros) == 0 {
		fmt.Println("No hay libros registrado")
		return
	}

	fmt.Println("Lista de Libros:")
	for _, libro := range Libros {
		estado := "Disponible"
		if !libro.Disponible {
			estado = "Prestado"
		}
		fmt.Printf("- %s | %s | %s | %s | %d | %s\n",
			libro.ID, libro.Titulo, libro.Autor, libro.Genero, libro.Year, estado)
	}
}
