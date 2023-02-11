package models

type Reports struct {
	ID          int    `json:"id"`
	Reporter_id int    `json:"reporter_id"`
	Reported_id int    `json:"reported_id"`
	Text        string `json:"text"`
	CreateAt    int64  `json:"created_at"`
}
