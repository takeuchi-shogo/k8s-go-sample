package models

type UserSearchFilters struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	Gender   *string `json:"gender"`
	Location *string `json:"location"`

	Purpose         int  `json:"purpose"`
	HasIntroduction bool `json:"has_introduction"`
	// Basic inromation
	HeightID           int `json:"height_id"`
	BodyTypeID         int `json:"body_type_id"`
	BloodTypeID        int `json:"blood_type_id"`
	ResidenceCountryID int `json:"residence_country_id"`
	ResidenceStateID   int `json:"residence_state_id"`
	HometownCountryID  int `json:"hometown_country_id"`
	HometownStateID    int `json:"hometown_state_id"`
	OccupationID       int `json:"occupation_id"`
	EducationID        int `json:"education_id"`
	AnnualIncomeID     int `json:"annual_income_id"`
	SmokingID          int `json:"smoking_id"`
	DrinkingID         int `json:"drinking_id"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
