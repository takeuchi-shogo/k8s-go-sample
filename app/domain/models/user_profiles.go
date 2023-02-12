package models

type UserProfiles struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	Interests   string `json:"interests"`
	LookingFor  string `json:"looking_for"`
	CreateAt    int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
