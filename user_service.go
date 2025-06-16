package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joesolano/libros-gestion/models"
)

var Usuarios []models.User

func RegisterUser() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')

	fmt.Print("correo: ")
	correo, _ := reader.ReadString('\n')

	fmt.Print("Rol (admin/lector): ")
	rol, _ := reader.ReadString('\n')

	id := fmt.Sprintf("U%d", time.Now().UnixNano())

	usuario := models.User{
		ID: id, 
		Nombre: strings.TrimSpace(nombre),
		Correo: strings.TrimSpace(correo),
		Rol: strings.ToLower(strings.TrimSpace(rol)),
	}

	Usuarios = append(Usuarios,usuario)
	fmt.Println("Usuario registrado con exito")

}

func ListUsers()  {
	if len(Usuarios) == 0 {
		fmt.Println("No hay usuarios registrados.")
		return
	}

	fmt.Println("Lista de Usarios: ") 
	for _, u := range Usuarios {
		fmt.Printf("- %s | %s | %s | %s\n", u.ID, u.Nombre, u.Correo, u.Rol)
	}
}
