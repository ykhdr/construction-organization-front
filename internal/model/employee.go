package model

type Employee struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Surname                string `json:"surname"`
	Patronymic             string `json:"patronymic"`
	Age                    int    `json:"age"`
	Seniority              int    `json:"seniority"`
	BuildingOrganizationID int    `json:"building_organization_id"`
}
