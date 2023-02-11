package models

type Users struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	ScreenName  string `json:"screen_name"`
	CreateAt    int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
