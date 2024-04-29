package model

type EngineerWorker struct {
	ID                     int              `json:"id"`
	SkillLevel             string           `json:"skill_level"`
	Position               EngineerPosition `json:"position"`
	TeamID                 int              `json:"team_id"`
	Name                   string           `json:"name"`
	Surname                string           `json:"surname"`
	Patronymic             string           `json:"patronymic"`
	Age                    int              `json:"age"`
	Seniority              int              `json:"seniority"`
	BuildingOrganizationID int              `json:"building_organization_id"`
}
