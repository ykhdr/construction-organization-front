package model

type MaterialUsage struct {
	EstimateID   int      `json:"estimate_id"`
	Material     Material `json:"material"`
	PlanQuantity int      `json:"plan_quantity"`
	FactQuantity int      `json:"fact_quantity"`
}
