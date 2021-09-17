package model

type User struct {
	Id        int    `json:"id"`
	Dpi       int    `json:"dpi"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
}

type UserR struct {
	Dpi       int    `json:"dpi"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
}

type Vaccine struct {
	Dosis      int `json:"dosis"`
	Vacuna_id  int `json:"vacuna_id"`
	Persona_id int `json:"persona_id"`
}

type Dosis struct {
	Id        int    `json:"id"`
	Dpi       int    `json:"dpi"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Vacuna    string `json:"vacuna"`
	Dosis     int    `json:"dosis"`
}
