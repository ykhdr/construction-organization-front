package model

import "time"

type Estimate struct {
	ID             int                 `json:"id"`
	MaterialUsage  []*MaterialUsageFit `json:"material_usage"`
	LastUpdateDate time.Time           `json:"last_update_date"`
}
