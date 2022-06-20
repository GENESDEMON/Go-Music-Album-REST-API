package main

import (
	"fmt"
	"golang-music-api/controllers"
	"golang-music-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//To load the configurations from the config.json using viper
	ConfigLoader()
	//Initialize the Database
	database.Connect(AppConfig.ConnectString)
	database.Migrate()
	//Initialize the Router
	router := mux.NewRouter().StrictSlash(true)
	//Register Album routes
	RegisterAlbumRoutes(router)

	//Start the server
	log.Println(fmt.Sprintf("Server starting on %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}

func RegisterAlbumRoutes(router *mux.Router) {
	router.HandleFunc("/api/albums", controllers.GetAlbums).Methods("GET")
	router.HandleFunc("/api/albums/{id}", controllers.GetAlbumById).Methods("GET")
	router.HandleFunc("/api/albums", controllers.CreateAlbum).Methods("POST")
	router.HandleFunc("/api/albums/{id}", controllers.UpdateAlbum).Methods("PUT")
	router.HandleFunc("/api/albums/{id}s", controllers.DeleteAlbum).Methods("DELETE")

}
