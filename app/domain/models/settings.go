package models

type Settings struct {
	ID                     int    `json:"id"`
	UserID                 int    `json:"user_id"`
	AllowEmailNotification bool   `json:"allow_email_notification"`
	AllowPushNotification  bool   `json:"allow_push_notification"`
	ContryCode             string `json:"contry_code"`
	Languege               string `json:"languege"`
	CreateAt               int64  `json:"created_at"`
	UpdatedAt              int64  `json:"updated_at"`
}
