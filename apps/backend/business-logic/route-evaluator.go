package grid

import (
	sql "github.com/JEFFTheDev/golang-exp/sql"
)

func GetShortestRoute(from string, to string) *[]string {
	shortestRoute := []string{}
	//currentDistance := math.MaxInt32
	//maxRoutesToCalc := 50

	connectedStationsAtStart := sql.GetConnectedTrainStationsCodesOfTrainStation(from)

	for _, st := range *connectedStationsAtStart {

		// First get the connected stations to the station
		connectedStationsAtStart = sql.GetConnectedTrainStationsCodesOfTrainStation(st)
	}

	return &shortestRoute
}
