package model

type BuildingSite struct {
	ID           int    `json:"id"`
	Address      string `json:"address"`
	ManagementID int    `json:"management_id"`
	ManagerID    int    `json:"manager_id"`
}
