package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

// clean_hotspots removes all non-NSW hotspot entries from the GeoJSON file retrieved from
// https://hotspots.dea.ga.gov.au/geoserver/wfs?service=wfs&version=2.0.0&request=GetFeature&outputFormat=application/json&typeNames=public:hotspots_three_days

// GeoJSON represents the GeoJSON format
type GeoJSON struct {
	Type           string     `json:"type"`
	Features       []Features `json:"features"`
	TotalFeatures  int        `json:"totalFeatures"`
	NumberMatched  int        `json:"numberMatched"`
	NumberReturned int        `json:"numberReturned"`
	TimeStamp      string     `json:"timeStamp"`
	Crs            Crs        `json:"crs"`
}

// Geometry fields of GeoJSON
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Properties fields of GeoJSON
type Properties struct {
	ID                       int         `json:"id"`
	Satellite                string      `json:"satellite"`
	SatelliteNssdcID         string      `json:"satellite_nssdc_id"`
	SatelliteOperatingAgency string      `json:"satellite_operating_agency"`
	Sensor                   string      `json:"sensor"`
	Orbit                    int         `json:"orbit"`
	StartDt                  string      `json:"start_dt"`
	StopDt                   string      `json:"stop_dt"`
	Filename                 string      `json:"filename"`
	ProcessDt                string      `json:"process_dt"`
	ProcessAlgorithm         string      `json:"process_algorithm"`
	ProcessAlgorithmVersion  string      `json:"process_algorithm_version"`
	Product                  string      `json:"product"`
	LoadDt                   string      `json:"load_dt"`
	Latitude                 float64     `json:"latitude"`
	Longitude                float64     `json:"longitude"`
	TempKelvin               float64     `json:"temp_kelvin"`
	Datetime                 string      `json:"datetime"`
	Power                    float64     `json:"power"`
	Confidence               int         `json:"confidence"`
	AustralianState          string      `json:"australian_state"`
	FireCategoryName         interface{} `json:"fire_category_name"`
	HoursSinceHotspot        float64     `json:"hours_since_hotspot"`
}

// Features fields of GeoJSON
type Features struct {
	Type         string     `json:"type"`
	ID           string     `json:"id"`
	Geometry     Geometry   `json:"geometry"`
	GeometryName string     `json:"geometry_name"`
	Properties   Properties `json:"properties"`
}

// CrsProperties fields of GeoJSON
type CrsProperties struct {
	Name string `json:"name"`
}

// Crs fields of GeoJSON
type Crs struct {
	Type       string        `json:"type"`
	Properties CrsProperties `json:"properties"`
}

func main() {
	var input = flag.String("input", "input.json", "JSON data to process")
	var output = flag.String("output", "output.json", "JSON data to ouput")

	flag.Parse()

	data, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatalf("could not open file for input: %v\n", err)
	}
	j := new(GeoJSON)
	err = json.Unmarshal(data, j)
	if err != nil {
		log.Fatalf("could not process JSON data: %v\n", err)
	}
	fs := make([]Features, 0)
	for _, f := range j.Features {
		if f.Properties.AustralianState == "NSW" && (f.Properties.TempKelvin <= 24.0 && f.Properties.TempKelvin > 0.0) {
			fs = append(fs, f)
		}
	}
	j.Features = fs
	out, err := json.Marshal(j)
	if err != nil {
		log.Fatalf("could not convert JSON data: %v\n", err)
	}
	err = ioutil.WriteFile(*output, out, 0644)
	if err != nil {
		log.Fatalf("could not write to file %s: %v\n", *output, err)
	}
}
