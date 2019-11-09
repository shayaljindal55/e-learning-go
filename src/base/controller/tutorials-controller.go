package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	db            *sql.DB
	err           error
	authenticated = false
)

// models
type Tutorials struct {
	Id          int    `json:"id"`
	Info        string `json:"info"`
	Description string `json:"description"`
}

func GetAllTutorials(w http.ResponseWriter, r *http.Request) {
	// check http method type
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	} else {
		database, _ :=
			sql.Open("sqlite3", "./bogo.db")
		rows, err :=
			database.Query("SELECT id, info, description FROM tutorials")
		if err != nil {
			fmt.Println("Prepare select all query error")
			panic(err)
		}
		var tutorials []Tutorials
		var tutorial Tutorials
		for rows.Next() {
			rows.Scan(&tutorial.Id, &tutorial.Info,
				&tutorial.Description)
			tutorials = append(tutorials, tutorial)
		}
		json.NewEncoder(w).Encode(tutorials)
	}
}

func AddTutorial(w http.ResponseWriter, r *http.Request) {
	// check http method type
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	} else {
		// parse the request body to read the inputs
		var t Tutorials
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)

		// Insert into database
		database, _ :=
			sql.Open("sqlite3", "./bogo.db")
		statement, err :=
			database.Prepare("INSERT INTO tutorials (info, description) VALUES (?, ?)")
		if err != nil {
			fmt.Println("Prepare insert query error")
			panic(err)
		}
		_, err = statement.Exec(t.Info, t.Description)
		if err != nil {
			fmt.Println("Execute insert query error")
			panic(err)
		}
		w.Write([]byte(`{"message": "Tutorial successfully inserted!"}`))
	}
}

func UpdateTutorial(w http.ResponseWriter, r *http.Request) {
	// check http method type
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	} else {
		// parse the request body to read the inputs
		var t Tutorials
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)

		// Update record in database
		database, err :=
			sql.Open("sqlite3", "./bogo.db")
		statement, err := database.Prepare(`
	UPDATE tutorials SET info=?, description=?
	WHERE id=?
`)
		if err != nil {
			fmt.Println("Prepare update query error")
			panic(err)
		}
		_, err = statement.Exec(t.Info, t.Description, t.Id)
		if err != nil {
			fmt.Println("Execute update query error")
			panic(err)
		}
		w.Write([]byte(`{"message": "Tutorial successfully updated!"}`))
	}
}

func DeleteTutorial(w http.ResponseWriter, r *http.Request) {
	// check http method type
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	} else {
		// parse the request body to read the inputs
		var t Tutorials
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)

		// Delete record from database
		database, err :=
			sql.Open("sqlite3", "./bogo.db")
		statement, err := database.Prepare("DELETE FROM tutorials WHERE id=?")
		if err != nil {
			fmt.Println("Prepare delete query error")
			panic(err)
		}
		_, err = statement.Exec(t.Id)
		if err != nil {
			fmt.Println("Execute delete query error")
			panic(err)
		}
		w.Write([]byte(`{"message": "Tutorial successfully deleted!"}`))
	}
}
