package model

type MasterTeam struct {
	ID                    int    `json:"id"`
	DesignExperienceYears int    `json:"design_experience_years"`
	Name                  string `json:"name"`
	ProjectID             int    `json:"project_id"`
}
