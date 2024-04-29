package model

type School struct {
	ID             int    `json:"id"`
	ClassroomCount int    `json:"classroom_count"`
	Floors         int    `json:"floors"`
	Name           string `json:"name"`
	BuildingSiteID int    `json:"building_site_id"`
	ProjectType    string `json:"type"`
}
