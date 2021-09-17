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
