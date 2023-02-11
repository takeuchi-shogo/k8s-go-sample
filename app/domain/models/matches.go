package models

type Matches struct {
	ID           int   `json:"id"`
	MaleUserID   int   `json:"male_user_id"`
	FemaleUserID int   `json:"female_user_id"`
	CreateAt     int64 `json:"created_at"`
}
