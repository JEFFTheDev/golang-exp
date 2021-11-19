package models

type TrackBehaviour interface {
}

type Track struct {
	Id          int    `json:"id"`
	Source      string `json:"source_station"`
	Destination string `json:"dest_station"`
	Cost        int    `json:"cost"`
}
