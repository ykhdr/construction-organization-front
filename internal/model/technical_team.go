package model

type TechnicalTeam struct {
	ID                  int    `json:"id"`
	IsMaintainMachinery bool   `json:"is_maintain_machinery"`
	Name                string `json:"name"`
	ProjectID           int    `json:"project_id"`
}
