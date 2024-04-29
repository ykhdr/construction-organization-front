package model

type ConstructionMachineryUsage struct {
	MachineryID    int `json:"machinery_id"`
	WorkScheduleID int `json:"work_schedule_id"`
	Quantity       int `json:"quantity"`
}
