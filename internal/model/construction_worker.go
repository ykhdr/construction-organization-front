package model

type ConstructionWorker struct {
	ID                     int    `json:"id"`
	IsShiftWorker          bool   `json:"is_shift_worker"`
	TeamID                 int    `json:"team_id"`
	Name                   string `json:"name"`
	Surname                string `json:"surname"`
	Patronymic             string `json:"patronymic"`
	Age                    int    `json:"age"`
	Seniority              int    `json:"seniority"`
	BuildingOrganizationID int    `json:"building_organization_id"`
}
