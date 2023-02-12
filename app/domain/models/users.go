package models

type Users struct {
	ID               int    `json:"id"`
	AccountID        int    `json:"account_id"`
	DisplayName      string `json:"display_name"`
	ScreenName       string `json:"screen_name"`
	Gender           string `json:"gender"`
	Location         string `json:"location"`
	IsAuthorizeEmail bool   `json:"is_authorize_email"`
	IsVerifiedAge    bool   `json:"is_verifed_age"`
	CreateAt         int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
	DeletedAt        *int64 `json:"deleted_at"`
}
