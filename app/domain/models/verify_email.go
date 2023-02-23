package models

type VerifyEmails struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	PINCode   string `json:"code"`
	ExpireAt  int64  `json:"expire_at"`
	CreatedAt int64  `json:"created_at"`
}
