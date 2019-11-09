package main

import (
	"fmt"
	"net/http"
	"os"
	"base/controller"
	_ "github.com/mattn/go-sqlite3"
)

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func main() {
	//	route
	http.HandleFunc("/getAllTutorials", controller.GetAllTutorials)
	http.HandleFunc("/addTutorial", controller.AddTutorial)
	http.HandleFunc("/updateTutorial", controller.UpdateTutorial)
	http.HandleFunc("/deleteTutorial", controller.DeleteTutorial)
	http.ListenAndServe(GetPort(), nil)
}