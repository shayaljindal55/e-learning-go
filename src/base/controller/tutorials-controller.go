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
	Url         string `json:"url"`
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func GetAllTutorials(w http.ResponseWriter, r *http.Request) {
	// check http method type
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	} else {
		database, _ :=
			sql.Open("sqlite3", "./bogo.db")
		rows, err :=
			database.Query("SELECT id, info, description, url FROM tutorials")
		if err != nil {
			fmt.Println("Prepare select all query error")
			panic(err)
		}
		var tutorials []Tutorials
		var t Tutorials
		for rows.Next() {
			rows.Scan(&t.Id, &t.Info,
				&t.Description, &t.Url)
			tutorials = append(tutorials, t)
		}
		resp := Message(true, "success")
		resp["tutorials"] = tutorials
		json.NewEncoder(w).Encode(resp)
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
			database.Prepare("INSERT INTO tutorials (info, description, url) VALUES (?, ?, ?)")
		if err != nil {
			fmt.Println("Prepare insert query error")
			panic(err)
		}
		_, err = statement.Exec(t.Info, t.Description, t.Url)
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
		fmt.Println(t.Url)

		// Update record in database
		database, err :=
			sql.Open("sqlite3", "./bogo.db")
		statement, err := database.Prepare(`
	UPDATE tutorials SET info=?, description=?, url=?
	WHERE id=?
`)
		if err != nil {
			fmt.Println("Prepare update query error")
			panic(err)
		}
		_, err = statement.Exec(t.Info, t.Description, t.Url, t.Id)
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
