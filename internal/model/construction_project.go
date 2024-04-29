package model

type ConstructionProject struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BuildingSiteID int    `json:"building_site_id"`
	ProjectType    string `json:"type"`
}
