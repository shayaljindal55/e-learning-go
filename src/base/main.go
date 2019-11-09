package main

import (
	"fmt"
	"net/http"
	"os"
	"base/controller"
	_ "github.com/mattn/go-sqlite3"
	 "github.com/rs/cors"
)

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
	 mux := http.NewServeMux()
	 handler := cors.Default().Handler(mux)
	mux.HandleFunc("/getAllTutorials", controller.GetAllTutorials)
	mux.HandleFunc("/addTutorial", controller.AddTutorial)
	mux.HandleFunc("/updateTutorial", controller.UpdateTutorial)
	mux.HandleFunc("/deleteTutorial", controller.DeleteTutorial)
	http.ListenAndServe(GetPort(), handler)
}