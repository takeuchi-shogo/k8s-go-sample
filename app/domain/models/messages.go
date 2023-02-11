package models

type Messages struct {
	ID         int    `json:"id"`
	MatchID    int    `json:"match_id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Message    string `json:"message"`
	CreateAt   int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	DeletedAt  *int64 `json:"deleted_at"`
}
