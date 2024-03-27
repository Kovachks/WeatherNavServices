package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

// getDirections is a function which pings the Google Maps API and returns a route between two points
func getDirections(r *http.Request) ([]maps.Route, error) {
	origin := r.URL.Query().Get("origin")
	destination := r.URL.Query().Get("destination")

	godotenv.Load()
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	results := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
	}
	route, _, err := c.Directions(context.Background(), results)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return route, err
}

func getWeatherResults(directions []maps.Route) ([]string, error) {
	// get the weather data
	return nil, nil
}

// formatWeatherDirectionData is a handler function that takes a request and response writer and returns a formatted JSON response
func formatWeatherDirectionData(w http.ResponseWriter, r *http.Request) {
	directionResults, err := getDirections(r)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	weatherResults, err := getWeatherResults(directionResults)

	var _ = weatherResults
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(directionResults)
}

func main() {
	http.HandleFunc("/directions", formatWeatherDirectionData)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	} else {
		log.Print("server started on port 3333")
	}

}
