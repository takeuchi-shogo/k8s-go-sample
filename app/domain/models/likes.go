package models

type Likes struct {
	ID            int   `json:"id"`
	SendUserID    int   `json:"send_user_id"`
	ReceiveUserID int   `json:"receive_user_id"`
	CreatedAt     int64 `json:"created_at"`
}
