package sql

import (
	"database/sql"
	"fmt"

	models "github.com/JEFFTheDev/golang-exp/models"
	_ "github.com/mattn/go-sqlite3"
)

func GetConnectedTrainStationsCodesOfTrainStation(code string) *[]string {
	var sqlResult = executeQuery("SELECT dest_station FROM Tracks WHERE source_station=?", code)

	var tracks []string

	for sqlResult.Next() {

		var track string
		var err = sqlResult.Scan(&track)
		if err != nil {
			fmt.Println("Error occurred: " + err.Error())
			return nil
		}

		tracks = append(tracks, track)
	}

	sqlResult.Close()

	return &tracks
}

func GetAllTrainStations() []models.TrainStation {
	var sqlResult = executeQuery("SELECT code, name_long, geo_lat, geo_lng FROM TrainStations", "")

	var stations []models.TrainStation

	for sqlResult.Next() {
		var station models.TrainStation
		var err = sqlResult.Scan(&station.Code, &station.Name, &station.Lat, &station.Lng)
		if err != nil {
			fmt.Println("Error occurred: " + err.Error())
			return nil
		}

		stations = append(stations, station)
	}

	sqlResult.Close()

	return stations
}

func GetAllTracks() *[]models.Track {
	var sqlResult = executeQuery("SELECT id, source_station, dest_station, cost FROM Tracks", "")

	var tracks []models.Track

	for sqlResult.Next() {

		var track models.Track
		var err = sqlResult.Scan(&track.Id, &track.Source, &track.Destination, &track.Cost)
		if err != nil {
			fmt.Println("Error occurred: " + err.Error())
			return nil
		}

		tracks = append(tracks, track)
	}

	sqlResult.Close()

	return &tracks
}

func executeQuery(query string, args string) *sql.Rows {
	db, err := sql.Open("sqlite3", "./infrastructure.db")
	if err != nil {
		fmt.Println("Error occurred: " + err.Error())
		return nil
	}

	var rows *sql.Rows

	if args == "" {
		rows, err = db.Query(query)
	} else {
		rows, err = db.Query(query, args)
	}

	if err != nil {
		fmt.Println("Error occurred: " + err.Error())
		return nil
	}

	db.Close()

	return rows
}
