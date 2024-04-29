package model

import "time"

type ConstructionContract struct {
	ID                     int       `json:"id"`
	BuildingOrganizationID int       `json:"building_organization_id"`
	CustomerOrganizationID int       `json:"customer_organization_id"`
	Name                   string    `json:"name"`
	ProjectID              int       `json:"project_id"`
	SigningDate            time.Time `json:"signing_date"`
}
