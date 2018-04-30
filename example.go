package main

import (
	apiClient "bandsintown-api-client/src"
	"fmt"
)

func main() {
	client := apiClient.NewClient()
	client.SetAppID("r@nd0m$tring")
	res, err := client.GetArtistInfo("Kanye West")
	fmt.Println(res, err)

	res2, err := client.GetEventsForArtist("Drake", nil, nil)
	fmt.Println(res2, err)
}
