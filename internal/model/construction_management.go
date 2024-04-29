package model

type ConstructionManagement struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ManagerID int    `json:"manager_id"`
}
