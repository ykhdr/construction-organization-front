package model

type EngineerTeam struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
	Type      string `json:"type"`
}
