package domain

type Child struct {
	ID       int    `json:"id"`       // ID del hijo
	Name     string `json:"name"`     // Nombre del hijo
	LastName string `json:"lastName"` // Apellido del hijo
	UserID   int    `json:"userId"`   // ID del usuario (padre/madre)
}
