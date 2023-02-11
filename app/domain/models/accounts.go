package models

type Accounts struct {
	ID          int    `json:"id"`
	User_ID     int    `json:"user_id"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreateAt    int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   *int64 `json:"deleted_at"`
}
