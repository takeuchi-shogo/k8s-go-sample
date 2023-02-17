package models

type Accounts struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	LoginStatus string `json:"login_status"`
	AccessLevel int    `json:"access_level"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   *int64 `json:"deleted_at"`
}
