package models

type TrainStationBehaviour interface {
}

type TrainStation struct {
	Id                int      `json:"id"`
	Code              string   `json:"code"`
	Name              string   `json:"name"`
	Lat               float64  `json:"lat"`
	Lng               float64  `json:"lng"`
	ConnectedStations []string `json:"connectedStations"`
}
