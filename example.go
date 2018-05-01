package main

import (
	apiClient "bandsintown-api-client/src"
	"fmt"
	"sort"
	"math"
	"strconv"
)

func main() {
	client := apiClient.NewClient()
	res, err := client.GetArtistInfo("Kanye West")
	fmt.Println(res, err)

	res2, err := client.GetEventsForArtist("30 Seconds To Mars", nil, nil)
	prettyPrintEvents(res2)

	sortEventsByGeo(res2, 60.0, 30.0)
	prettyPrintEvents(res2)
}

func sortEventsByGeo(events []*apiClient.Event, lat, long float64) {
	sort.Slice(events, func(i, j int) bool {
		lat1, _ := strconv.ParseFloat(events[i].Venue.Latitude, 64)
		long1, _ := strconv.ParseFloat(events[i].Venue.Longitude, 64)

		lat2, _ := strconv.ParseFloat(events[j].Venue.Latitude, 64)
		long2, _ := strconv.ParseFloat(events[j].Venue.Longitude, 64)

		distance1 := math.Pow(lat1 - lat, 2) + math.Pow(long1 - long, 2)
		distance2 := math.Pow(lat2 - lat, 2) + math.Pow(long2 - long, 2)

		return distance1 < distance2
	})
}

func prettyPrintEvents(events []*apiClient.Event) {
	for _, event := range events {
		fmt.Println(event.Id + " " + event.Datetime + " " + event.Venue.City)
	}

	fmt.Println("\n")
}
