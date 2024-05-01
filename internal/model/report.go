package model

import (
	"time"
)

type Report struct {
	ID                 int        `json:"id"`
	ProjectID          int        `json:"project_id"`
	ReportCreationDate time.Time  `json:"report_creation_date"`
	ReportFile         ReportFile `json:"report_file"`
}
