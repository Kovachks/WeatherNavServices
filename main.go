package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func getDirections(w http.ResponseWriter, r *http.Request) {
	pretty.Log(r.URL.Query())
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

	return route
}

func formatWeatherDirectionData(w http.ResponseWriter, r *http.Request) {
	getDirectionsFunc := getDirections
	directionResults := getDirectionsFunc(w, r)

	log.Printf(directionResults)

}

func main() {
	http.HandleFunc("/directions", formatWeatherDirectionData)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
