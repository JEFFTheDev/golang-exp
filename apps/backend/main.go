package main

import (
	"encoding/json"
	"log"
	"net/http"

	logic "github.com/JEFFTheDev/golang-exp/business-logic"
	parser "github.com/JEFFTheDev/golang-exp/track-data-parser"
)

func main() {
	parser.ReadJson()
	handleRequests()
}

func getTrainStations(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var gridDb = logic.GetTrainStations()
	json.NewEncoder(w).Encode(gridDb)
}

func getTracks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var gridDb = logic.GetTracks()
	json.NewEncoder(w).Encode(gridDb)
}

func handleRequests() {
	http.HandleFunc("/api/gettrainstations", getTrainStations)
	http.HandleFunc("/api/gettracks", getTracks)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
