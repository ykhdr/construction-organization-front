package model

import "time"

type Estimate struct {
	ID             int              `json:"id"`
	MaterialUsage  []*MaterialUsage `json:"material_usage"`
	LastUpdateDate time.Time        `json:"last_update_date"`
}
