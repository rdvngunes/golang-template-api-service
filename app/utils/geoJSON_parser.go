package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

type GeoJSON struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type WaypointsGeoJSON struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type Properties struct {
	Title string `json:"title"`
}

var CVProperties = `{"fill": "yellow", "title": "Contingency Volume"}`
var GRBProperties = `{"fill": "red", "title": "Ground Risk Buffer"}`
var OperationalVolumeProperties = `{"title": "Operational Volume"}`

func GeoJSONFetaureExtractor(geojson string, inputField string) (string, error) {
	fmt.Println("Inside GeoJSON extractor.....")
	//inputField = "geometry"
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(geojson), &data); err != nil {
		fmt.Println("Error parsing input string:", err)
		return "", err
	}

	geoJSONData := data["features"].([]interface{})[0].(map[string]interface{})[inputField]

	// Convert GeoJSON data to string
	geoJSONString, err := json.Marshal(geoJSONData)
	if err != nil {
		fmt.Println("Error marshaling GeoJSON data:", err)
		return "", err
	}

	return string(geoJSONString), nil
}

func GeoJSONToPolygonFormat(geoJSONString string, multiPolygon bool) (string, error) {
	var formattedCoordinates []string

	var geoJSON GeoJSON
	err := json.Unmarshal([]byte(geoJSONString), &geoJSON)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	for _, ring := range geoJSON.Coordinates {
		var points []string
		for _, point := range ring {
			// Only include first two elements (latitude and longitude) in the formatted point
			points = append(points, fmt.Sprintf("%.6f %.6f", point[0], point[1]))
		}
		formattedRing := strings.Join(points, ", ")
		formattedCoordinates = append(formattedCoordinates, fmt.Sprintf("(%s)", formattedRing))
	}
	if multiPolygon {
		return fmt.Sprintf("MULTIPOLYGON((%s))", strings.Join(formattedCoordinates, ", ")), nil

	} else {
		return fmt.Sprintf("POLYGON((%s))", strings.Join(formattedCoordinates, ", ")), nil
	}
}

func WaypointsParser(data string) (error, *string) {
	var geo WaypointsGeoJSON
	err := json.Unmarshal([]byte(data), &geo)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return err, nil
	}

	// Filter out features based on the title in properties
	var filteredFeatures []Feature
	for _, feature := range geo.Features {
		if feature.Properties.Title != "Contingency Volume" && feature.Properties.Title != "Ground Risk Buffer" {
			filteredFeatures = append(filteredFeatures, feature)
		}
	}

	// Update the GeoJSON with filtered features
	geo.Features = filteredFeatures

	// Convert back to JSON to output
	filteredJSON, err := json.MarshalIndent(geo, "", "    ")
	if err != nil {
		fmt.Println("Error generating JSON: ", err)
		return err, nil
	}
	WaypointsList := string(filteredJSON)

	fmt.Println(WaypointsList)
	return nil, &WaypointsList
}
