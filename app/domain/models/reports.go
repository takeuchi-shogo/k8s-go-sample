package models

type Reports struct {
	ID         int     `json:"id"`
	ReporterID int     `json:"reporter_id"`
	ReportedID int     `json:"reported_id"`
	Reason     *string `json:"reason"`
	CreateAt   int64   `json:"created_at"`
}
