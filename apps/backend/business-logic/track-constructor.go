package grid

import (
	models "github.com/JEFFTheDev/golang-exp/models"
	sql "github.com/JEFFTheDev/golang-exp/sql"
)

func GetTrainStations() *[]models.TrainStation {

	// Retrieve all the trainstations.
	var trainstations = sql.GetAllTrainStations()

	for i, station := range trainstations {
		var connectedStationCodes = sql.GetConnectedTrainStationsCodesOfTrainStation(station.Code)
		trainstations[i].ConnectedStations = *connectedStationCodes
	}

	return &trainstations
}

func GetTracks() *[]models.Track {
	return sql.GetAllTracks()
}
