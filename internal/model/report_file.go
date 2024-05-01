package model

import (
	"time"
)

type ReportFile struct {
	ID        int           `json:"id"`
	ProjectID int           `json:"project_id"`
	Schedules []ScheduleFit `json:"schedules"`
	Estimate  EstimateFit   `json:"estimate.html"`
}

type ScheduleFit struct {
	WorkType      WorkType
	Team          TeamFit   `json:"construction_team"`
	PlanStartDate time.Time `json:"plan_start_date"`
	PlanEndDate   time.Time `json:"plan_end_date"`
	FactStartDate time.Time `json:"fact_start_date"`
	FactEndDate   time.Time `json:"fact_end_date"`
	PlanOrder     int       `json:"plan_order"`
	FactOrder     int       `json:"fact_order"`
}

type TeamFit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EstimateFit struct {
	MaterialUsage  []MaterialUsageFit `json:"material_usage"`
	LastUpdateDate time.Time          `json:"last_update_date"`
}

type MaterialUsageFit struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Cost         int    `json:"cost"`
	PlanQuantity int    `json:"plan_quantity"`
	FactQuantity int    `json:"fact_quantity"`
}
