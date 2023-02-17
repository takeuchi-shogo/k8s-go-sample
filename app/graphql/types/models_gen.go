// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package types

type Matches struct {
	ID           string `json:"id"`
	MaleUserID   int    `json:"male_user_id"`
	FemaleUserID int    `json:"female_user_id"`
}

type NewAccounts struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type NewBlocks struct {
	Blocking int `json:"blocking"`
	Blocked  int `json:"blocked"`
}

type NewReports struct {
	ReporterID int    `json:"reporter_id"`
	ReportedID int    `json:"reported_id"`
	Reason     string `json:"reason"`
}

type NewUsers struct {
	DisplayName string `json:"display_name"`
	ScreenName  string `json:"screen_name"`
	Gender      string `json:"gender"`
}

type NewVerifyEmails struct {
	Email string `json:"email"`
}

type VerifyEmails struct {
	ID      string `json:"id"`
	Email   string `json:"Email"`
	Token   string `json:"Token"`
	PINCode string `json:"PINCode"`
}
