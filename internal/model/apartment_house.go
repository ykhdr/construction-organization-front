package model

type ApartmentHouse struct {
	ID             int    `json:"id"`
	Floors         int    `json:"floors"`
	Name           string `json:"name"`
	BuildingSiteID int    `json:"building_site_id"`
	ProjectType    string `json:"type"`
}
