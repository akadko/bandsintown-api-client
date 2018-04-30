package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"os"
	"fmt"
	"time"
	"io/ioutil"
	"strconv"
	"strings"
)

type Client struct {
	baseUrl string
	appId string

	http *http.Client
}

type ApiError struct {
	Message string `json:"message"`
	Errors []string `json:"errors"`
	StatusCode int
}

func (apiError ApiError) Error() string {
	return strconv.Itoa(apiError.StatusCode) + ": " + apiError.Message + "\n" + strings.Join(apiError.Errors, "\n")
}

func NewClient() *Client {
	c := Client{
		baseUrl: "https://rest.bandsintown.com/",
		appId: os.Getenv("BANDSINTOWN_APP_ID"),
		http: &http.Client{},
	}

	return &c
}

func (c *Client) Set(key string, value string) error {
	switch key {
	case "baseUrl":
		c.baseUrl = value
		return nil
	case "appId":
		c.appId = value
		return nil
	}

	return fmt.Errorf("Wrong parameter \"%s\" of Bandsintown API Client.", key)
}

func (c *Client) GetArtistInfo(artistName string) (*Artist, error) {
	apiUrl := c.baseUrl + "artists/" + url.PathEscape(artistName)
	v := url.Values{}
	v.Set("app_id", c.appId)
	apiUrl = apiUrl + "?" + v.Encode()

	resp, err := c.http.Get(apiUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError ApiError
		json.Unmarshal(bodyBytes, &apiError)
		apiError.StatusCode = resp.StatusCode

		return nil, apiError
	}

	var artist Artist
	err = json.Unmarshal(bodyBytes, &artist)

	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (c *Client) GetEventsForArtist(artistName string, dateStart *time.Time, dateEnd *time.Time) ([]Event, error) {
	apiUrl := c.baseUrl + "artists/" + url.PathEscape(artistName) + "/events"
	v := url.Values{}
	v.Set("app_id", c.appId)

	if dateStart != nil && dateEnd != nil {
		date := dateStart.Format("2017-12-31") + "," + dateEnd.Format("2017-12-31")
		v.Set("date", date)
	}

	apiUrl = apiUrl + "?" + v.Encode()
	resp, err := c.http.Get(apiUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError ApiError
		json.Unmarshal(bodyBytes, &apiError)
		apiError.StatusCode = resp.StatusCode

		return nil, apiError
	}

	if err != nil {
		return nil, err
	}

	var events []Event
	err = json.Unmarshal(bodyBytes, events)

	if err != nil {
		return nil, err
	}

	return events, nil
}
