package src

import (
	"net/http"
	"net/url"
	"encoding/json"
	"os"
	"time"
	"io/ioutil"
	"strconv"
	"strings"
)

type Client struct {
	baseURL string
	appId   string

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
		baseURL: "https://rest.bandsintown.com/",
		appId:   os.Getenv("BANDSINTOWN_APP_ID"),
		http:    &http.Client{},
	}

	return &c
}

func (c *Client) SetAppID(appId string) {
	c.appId = appId
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) SetHTTPClient(http *http.Client) {
	c.http = http
}

func (c *Client) GetArtistInfo(artistName string) (*Artist, error) {
	apiUrl := c.baseURL + "artists/" + url.PathEscape(artistName)
	v := url.Values{}
	v.Set("app_id", c.appId)
	apiUrl = apiUrl + "?" + v.Encode()

	var artist Artist
	err := c.doGetRequest(apiUrl, &artist)

	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (c *Client) GetEventsForArtist(artistName string, dateStart *time.Time, dateEnd *time.Time) ([]*Event, error) {
	apiUrl := c.baseURL + "artists/" + url.PathEscape(artistName) + "/events"
	v := url.Values{}
	v.Set("app_id", c.appId)

	if dateStart != nil && dateEnd != nil {
		date := dateStart.Format("2017-12-31") + "," + dateEnd.Format("2017-12-31")
		v.Set("date", date)
	}

	apiUrl = apiUrl + "?" + v.Encode()
	var events []*Event

	err := c.doGetRequest(apiUrl, &events)

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (c *Client) doGetRequest(url string, result interface{}) (error) {
	resp, err := c.http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError ApiError
		json.Unmarshal(bodyBytes, &apiError)
		apiError.StatusCode = resp.StatusCode

		return apiError
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, result)

	if err != nil {
		return err
	}

	return nil
}