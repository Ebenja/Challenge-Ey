package main

import (
	util "Golang-backend/Util"
	"Golang-backend/model"
	user "Golang-backend/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

func routeIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "API backend Golang Challenge")
}

func createRegister(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	db := util.InitDB()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}
	// insert into registro(nombre,apellido,dpi,ciudad,vacuna,fecha) values ('Juan','Lopez',111,'Chimaltenango','Moderna', STR_TO_DATE('28-07-2021','%d-%m-%Y') )

	sql := "insert into registro(nombre,apellido,dpi,ciudad,vacuna,fecha) values (?, ?, ?, ?, ?,  STR_TO_DATE( ? ,'%d-%m-%Y'))"

	// Ejeucion
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec("Juani", "Lopez", 111, "Chimaltenango", "Moderna", "28-07-202")
	if err != nil {
		panic(err)
	}
	tx.Commit()
	fmt.Println(result.LastInsertId())

}

func allUsers(w http.ResponseWriter, req *http.Request) {
	print("ebtre")
	db := util.InitDB()
	db.Begin()

	rows, err := db.Query("select  * from persona")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listaTemps []user.User
	for rows.Next() {
		var temps user.User
		err = rows.Scan(&temps.Id, &temps.Dpi, &temps.Nombre, &temps.Apellidos)
		if err != nil {
			panic(err.Error())
		}

		listaTemps = append(listaTemps, temps)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listaTemps)
	db.Close()

}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user.UserR
	db := util.InitDB()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	// se envia
	//  {
	// 	"dpi" : 1234,
	// 	"nombre" : "Benjamin",
	// 	"apellidos" : "Lopez"
	// }
	// currentTime := time.Now()
	rows, err := db.Query("INSERT INTO persona (dpi,nombre,apellidos) VALUES (?,?,?)", &newUser.Dpi, &newUser.Nombre, &newUser.Apellidos)
	// rows, err := db.Query("INSERT INTO usuario (username,passwordd,nombre,apellido,fecha_nacimiento,fecha_registro,email,foto) VALUES ('" + newUser.Username + "','" + newUser.Password + "','" + newUser.Nombre + "','" + newUser.Apellido + "',TO_DATE('" + newUser.Fecha_nacimiento + "','dd/mm/yyyy'),TO_DATE('" + currentTime.Format("02/01/2006") + "','dd/mm/yyyy'),'" + newUser.Email + "','" + path + "')")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
	db.Close()

}

func allVacccine(w http.ResponseWriter, req *http.Request) {
	print("ebtre")
	db := util.InitDB()
	db.Begin()

	// rows, err := db.Query("select p.*,v.vacuna,d.dosis from persona p inner join dosis d on d.persona_id=p.id inner join vacuna v on v.id=d.vacuna_id")
	// select p.*,v.vacuna,count(d.dosis) as dosis from persona p inner join dosis d on d.persona_id=p.id inner join vacuna v on v.id=d.vacuna_id group by v.vacuna
	rows, err := db.Query("select p.*,v.vacuna,count(d.dosis) as dosis from persona p inner join dosis d on d.persona_id=p.id inner join vacuna v on v.id=d.vacuna_id group by v.vacuna")

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	// print(rows)

	var listaTemps []user.Dosis
	for rows.Next() {
		var temps user.Dosis
		err = rows.Scan(&temps.Id, &temps.Dpi, &temps.Nombre, &temps.Apellidos, &temps.Vacuna, &temps.Dosis)
		if err != nil {
			panic(err.Error())
		}
		// print(temps)

		listaTemps = append(listaTemps, temps)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listaTemps)
	db.Close()

}

func allVacccineReport(w http.ResponseWriter, req *http.Request) {
	print("ebtre")
	db := util.InitDB()
	db.Begin()

	rows, err := db.Query("select p.*,v.vacuna,d.dosis from persona p inner join dosis d on d.persona_id=p.id inner join vacuna v on v.id=d.vacuna_id")
	// select p.*,v.vacuna,count(d.dosis) as dosis from persona p inner join dosis d on d.persona_id=p.id inner join vacuna v on v.id=d.vacuna_id group by v.vacuna
	// rows, err := db.Query("select p.*,v.vacuna,count(d.dosis) as dosis from persona p inner join dosis d on d.persona_id=p.id inner join vacuna v on v.id=d.vacuna_id group by v.vacuna")

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	// print(rows)

	var listaTemps []user.Dosis
	for rows.Next() {
		var temps user.Dosis
		err = rows.Scan(&temps.Id, &temps.Dpi, &temps.Nombre, &temps.Apellidos, &temps.Vacuna, &temps.Dosis)
		if err != nil {
			panic(err.Error())
		}
		// print(temps)

		listaTemps = append(listaTemps, temps)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listaTemps)
	db.Close()

}

func addVaccine(w http.ResponseWriter, r *http.Request) {
	var newUser user.Vaccine
	db := util.InitDB()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	// se envia
	//  {
	// 	"dpi" : 1234,
	// 	"nombre" : "Benjamin",
	// 	"apellidos" : "Lopez"
	// }
	// currentTime := time.Now()
	rows, err := db.Query("INSERT INTO dosis (dosis, fecha, vacuna_id, persona_id) VALUES (?, DATE(NOW()), ?, ?)", &newUser.Dosis, &newUser.Vacuna_id, &newUser.Persona_id)
	// rows, err := db.Query("INSERT INTO usuario (username,passwordd,nombre,apellido,fecha_nacimiento,fecha_registro,email,foto) VALUES ('" + newUser.Username + "','" + newUser.Password + "','" + newUser.Nombre + "','" + newUser.Apellido + "',TO_DATE('" + newUser.Fecha_nacimiento + "','dd/mm/yyyy'),TO_DATE('" + currentTime.Format("02/01/2006") + "','dd/mm/yyyy'),'" + newUser.Email + "','" + path + "')")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
	db.Close()

}

func getUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var newTemp model.User
	dpiuser := query.Get("dpiuser")
	println(dpiuser)
	db := util.InitDB()
	err := db.QueryRow("SELECT * from persona WHERE dpi=?", dpiuser).Scan(&newTemp.Id, &newTemp.Dpi, &newTemp.Nombre, &newTemp.Apellidos)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newTemp)
	db.Close()
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var newChange model.User
	// var newTemp model.User
	// var newChange2 model.User
	db := util.InitDB()
	reqBody, err := ioutil.ReadAll(r.Body)
	print("aqui ando")
	if err != nil {
		print("error ujpdate parameters")
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newChange)
	println("id: " + strconv.Itoa(newChange.Id))

	rows, err1 := db.Query("UPDATE persona SET dpi=?,  nombre=?,  apellidos=?  WHERE id=?", &newChange.Dpi, &newChange.Nombre, &newChange.Apellidos, &newChange.Id)
	if err1 != nil {
		// panic(err1.Error())
		print("error  sdfasd")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	json.NewEncoder(w).Encode(resp)
	db.Close()
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	db := util.InitDB()
	// reqBody, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Fprintf(w, "Insert a Valid User Data")
	// }
	query := r.URL.Query()
	dpiuser := query.Get("dpi")
	println(dpiuser)

	// json.Unmarshal(reqBody, &newUser)
	rows, err := db.Query("DELETE FROM persona WHERE dpi = " + dpiuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
	db.Close()
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", routeIndex)
	router.HandleFunc("/User/getAll", allUsers).Methods("GET")
	router.HandleFunc("/User/create", createUser).Methods("POST")
	router.HandleFunc("/User/get", getUser).Methods("GET")
	router.HandleFunc("/User/update", updateUser).Methods("PUT")
	router.HandleFunc("/User/delete", deleteUser).Methods("DELETE")

	router.HandleFunc("/Dosis/add", addVaccine).Methods("POST")
	router.HandleFunc("/Dosis/report", allVacccine).Methods("GET")
	router.HandleFunc("/Dosis/All", allVacccineReport).Methods("GET")

	http.ListenAndServe(":3000", router)

	log.Println("esta corriendo el servidor Port:3000")

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
