package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JEFFTheDev/golang-exp/grid"
)

func main() {

	handleRequests()
}

func getGrid(w http.ResponseWriter, r *http.Request) {
	var gridDb = grid.GetGridWithNodesAndEdgesById(2)
	json.NewEncoder(w).Encode(gridDb)
}

func handleRequests() {
	http.HandleFunc("/grid", getGrid)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
