package models

type Users struct {
	ID               int    `json:"id"`
	UUID             string `json:"uuid"`
	AccountID        int    `json:"account_id"`
	DisplayName      string `json:"display_name"`
	ScreenName       string `json:"screen_name"`
	Age              int    `json:"age"`
	Gender           string `json:"gender"`
	Location         string `json:"location"`
	IsAuthorizeEmail bool   `json:"is_authorize_email"`
	IsVerifiedAge    bool   `json:"is_verified_age"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
	DeletedAt        *int64 `json:"deleted_at"`
}

type ResponseUsers struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	// AccountID        int    `json:"account_id"`
	DisplayName      string `json:"display_name"`
	ScreenName       string `json:"screen_name"`
	Age              int    `json:"age"`
	Gender           string `json:"gender"`
	Location         string `json:"location"`
	IsAuthorizeEmail bool   `json:"is_authorize_email"`
	IsVerifiedAge    bool   `json:"is_verified_age"`

	UserProfile *ResponseUserProfiles `json:"user_profile"`

	IsLiked bool `json:"is_liked"` // 既にいいねしたユーザーか

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
