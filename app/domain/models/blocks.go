package models

type Blocks struct {
	ID         int   `json:"id"`
	BlockingID int   `json:"blocking_id"`
	BlockedID  int   `json:"blocked_id"`
	CreateAt   int64 `json:"created_at"`
}
