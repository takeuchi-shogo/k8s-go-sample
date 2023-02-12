package models

type Likes struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	CreateAt   int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	DeletedAt  *int64 `json:"deleted_at"`
}
