package main

type Artist struct {
	Id int `json:"id"`
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
	ArtistId string `json:"artist_id"`
	Url string `json:"url"`
	OnSaleDatetime string `json:"on_sale_datetime"`
	Datetime string `json:"datetime"`
	Description string `json:"description"`
	Venue *Venue `json:"venue"`
	Offers []*Offer `json:"offers"`
	Lineup []string `json:"lineup"`
}

type Venue struct {
	Name string `json:"name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City string `json:"city"`
	Region string `json:"region"`
	Country string `json:"country"`
}

type Offer struct {
	OfferType string `json:"type"`
	Url string `json:"url"`
	Status string `json:"status"`
}