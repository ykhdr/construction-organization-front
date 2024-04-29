package model

type Bridge struct {
	ID                 int    `json:"id"`
	Span               int    `json:"span"`
	Width              int    `json:"width"`
	TrafficLanesNumber int    `json:"traffic_lanes_number"`
	Name               string `json:"name"`
	BuildingSiteID     int    `json:"building_site_id"`
	ProjectType        string `json:"type"`
}
