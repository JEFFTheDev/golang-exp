package track_data_parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadJson() {

	// First open the tracks.json file
	trackFile, err := os.Open("./track-data-parser/tracks.json")
	trackData, _ := ioutil.ReadAll(trackFile)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Opened tracks.json...")

	var tracks []map[string]interface{}

	// Parse de JSON naar een array met maps.
	json.Unmarshal(trackData, &tracks)

	sql := ""

	for _, v := range tracks {
		sourceStation := ""
		for key, value := range v {

			if key == "Station" {
				sourceStation = fmt.Sprintf("%v", value)
				continue
			}

			// XXX betekent dat het station hetzelfde is als het hoofdstation, hier kan niet naar verbonden worden.
			// ? betekent dat het station niet is verbonden.
			if value == "XXX" || value == "?" {
				continue
			}

			if sourceStation == "" {
				continue
			}

			// TODO: Waarom verschijnt station ZDK en aantal andere stations niet??
			sql += fmt.Sprintf(
				"\nINSERT INTO Tracks (source_station, dest_station, cost) VALUES( '%s', '%s', %s);",
				sourceStation, fmt.Sprintf("%v", key), fmt.Sprintf("%v", value))
		}
	}

	// Create the sql file to write to.
	f, err := os.Create("./track-data-parser/tracks.sql")

	defer f.Close()

	_, err2 := f.WriteString(sql)

	if err2 != nil {
		log.Fatal(err2)
	}

	defer trackFile.Close()
}
