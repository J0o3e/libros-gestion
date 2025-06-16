package models

import "time"

type Loan struct {
	ID string
	LibroID string
	UsuarioID string
	Inicio time.Time
	Fin *time.Time
	Estado string //En este caso puede ser prestado o devuelto// 
}
