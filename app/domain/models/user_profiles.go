package models

type UserProfiles struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	Introduction *string `json:"introduction"`
	Interests    *string `json:"interests"`
	LookingFor   *string `json:"looking_for"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
	DeletedAt    *int64  `json:"deleted_at"`
}
