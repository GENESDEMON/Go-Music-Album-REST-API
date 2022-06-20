package controllers

import (
	"encoding/json"
	"golang-music-api/database"
	"golang-music-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"application/json")
	var album models.Albums
	json.NewDecoder(r.Body).Decode(&album)
	database.Instance.Create(&album)
	json.NewEncoder(w).Encode(album)
}

func checkIfAlbumExists(albumId string) bool {
	var album models.Albums
	database.Instance.First(&album, albumId)
	if album.ID == 0 {
		return false
	}
	return true
}

func GetAlbumById(w http.ResponseWriter, r *http.Request) {
	albumId := mux.Vars(r)["id"]
	if checkIfAlbumExists(albumId) == false {
		json.NewEncoder(w).Encode("Album Not Found!")
		return
	}
	var album models.Albums
	database.Instance.First(&album, albumId)
	w.Header().Set(
		"Content-Type",
		"application/json")
	json.NewEncoder(w).Encode(album)
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	var albums []models.Albums
	database.Instance.Find(&albums)
	w.Header().Set(
		"Content-Type",
		"application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	albumId := mux.Vars(r)["id"]
	if checkIfAlbumExists(albumId) == false {
		json.NewEncoder(w).Encode("Album Not Found!")
		return
	}
	var album models.Albums
	database.Instance.First(&album, albumId)
	json.NewDecoder(r.Body).Decode(&album)
	database.Instance.Save(&album)
	w.Header().Set(
		"Content-Type",
		"application/json")
	json.NewEncoder(w).Encode(album)

}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"application/json")
	albumId := mux.Vars(r)["id"]
	if checkIfAlbumExists(albumId) == false {
		json.NewEncoder(w).Encode("Album Not Found!")
		return
	}
	var album models.Albums
	database.Instance.First(&album, albumId)
	json.NewEncoder(w).Encode("Alum Deleted Successfully!")
}
