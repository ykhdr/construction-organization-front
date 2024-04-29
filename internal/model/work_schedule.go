package model

import "time"

type WorkSchedule struct {
	ID                 int       `json:"id"`
	ConstructionTeamID int       `json:"construction_team_id"`
	WorkType           WorkType  `json:"work_type_id"`
	PlanStartDate      time.Time `json:"plan_start_date"`
	PlanEndDate        time.Time `json:"plan_end_date"`
	FactStartDate      time.Time `json:"fact_start_date"`
	FactEndDate        time.Time `json:"fact_end_date"`
	PlanOrder          int       `json:"plan_order"`
	FactOrder          int       `json:"fact_order"`
	ProjectID          int       `json:"project_id"`
}
