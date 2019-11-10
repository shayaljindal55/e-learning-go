package main

import (
	"fmt"
	"net/http"
	"os"
	"base/controller"
	_ "github.com/mattn/go-sqlite3"
	 "github.com/rs/cors"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	enableCors(&w)
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}
// }

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
//     (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//     (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

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