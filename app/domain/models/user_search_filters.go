package models

type UserSearchFilters struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Gender    *string `json:"gender"`
	Location  *string `json:"location"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}
