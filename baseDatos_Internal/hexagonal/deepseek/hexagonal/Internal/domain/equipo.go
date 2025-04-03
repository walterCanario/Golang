// internal/domain/equipo.go
package domain

type Equipo struct {
	ID      string `json:"id"`
	Nombre  string `json:"nombre"`
	Ciudad  string `json:"ciudad"`
	Estadio string `json:"estadio"`
}
