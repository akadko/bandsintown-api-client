package src

import (
	"time"
	"strings"
	"fmt"
)

type Artist struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	ImageUrl string `json:"image_url"`
	ThumbUrl string `json:"thumb_url"`
	FacebookPageURL string `json:"facebook_page_url"`
	MBID string `json:"mbid"`
	TrackerCount int `json:"tracker_count"`
	UpcomingEventCount int `json:"upcoming_event_count"`
}

type Event struct {
	Id string `json:"id"`
	ArtistName string
	ArtistId string `json:"artist_id"`
	Url string `json:"url"`
	OnSaleDatetime string `json:"on_sale_datetime"`
	Datetime EventTime `json:"datetime"`
	Description string `json:"description"`
	Venue *Venue `json:"venue"`
	Offers []*Offer `json:"offers"`
	Lineup []string `json:"lineup"`
}

type Venue struct {
	Name string `json:"name"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	City string `json:"city"`
	Region string `json:"region"`
	Country string `json:"country"`
}

type Offer struct {
	OfferType string `json:"type"`
	Url string `json:"url"`
	Status string `json:"status"`
}

type EventTime struct {
	time.Time
}

const etLayout = "2006-01-02T15:04:05"

func (et *EventTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		et.Time = time.Time{}
		return
	}
	et.Time, err = time.Parse(etLayout, s)
	return
}

func (et *EventTime) MarshalJSON() ([]byte, error) {
	if et.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", et.Time.Format(etLayout))), nil
}