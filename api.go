package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AthanatiusC/FaceRecogApi/controllers"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	// "github.com/zuramai/smartschool_api/models"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "socket-test.html")
		fmt.Println("home")
	})
	// http.Handle("/", http.FileServer(http.Dir("./assets")))
	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "8088"
	}

	apiV2 := router.PathPrefix("/api/v2").Subrouter()

	v2User := apiV2.PathPrefix("/user").Subrouter()
	v2User.HandleFunc("/recognize", controller.Recognize).Methods("OPTION", "POST")                     // Recognize
	
	fmt.Println("App running on port " + APP_PORT)
	log.Fatal(http.ListenAndServe(":"+APP_PORT, router))
}

