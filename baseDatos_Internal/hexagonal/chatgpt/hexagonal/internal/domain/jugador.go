package domain

type Jugador struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Equipo   string `json:"equipo"`
	Edad     int    `json:"edad"`
	Posicion string `json:"posicion"`
}
