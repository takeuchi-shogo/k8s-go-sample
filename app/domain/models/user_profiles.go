package models

type UserProfiles struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	Purpose      int     `json:"purpose"`
	Introduction *string `json:"introduction"`
	// Basic Profile
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
	// User Information
	// Nickname
	// Age
	SiblingsID   int `json:"siblings_id"`
	LanguageID   int `json:"language_id"`
	InterestsID  int `json:"interests_id"`
	LookingForID int `json:"looking_for_id"`
	// Education and Career
	SchoolName *string `json:"school_name"`
	JobTitle   *string `json:"job_title"`
	// Attitude towards Love and Marriage
	MaritalHistoryID                 int `json:"marital_history_id"`
	PresenceOfChildrenID             int `json:"presence_of_children_id"`
	IntentionsTowardsMarriageID      int `json:"intentions_towards_marriage_id"`
	DesireForChildrenID              int `json:"desire_for_children_id"`
	HouseholdChoresAndChildRearingID int `json:"household_chores_and_child_rearing_id"`
	MeetingPreferenceID              int `json:"meeting_preference_id"`
	DatingExpensesID                 int `json:"dating_expenses_id"`
	// Personality, Lifestyle, Hobbies
	PersonalityTypeID int `json:"presonality_type_id"`
	SociabilityID     int `json:"sociability_id"`
	RoommatesID       int `json:"roommates_id"`
	DaysOffID         int `json:"days_off_id"`
	HobbiesID         int `json:"hobbies_id"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"deleted_at"`
}

type ResponseUserProfiles struct {
	UserID       int     `json:"user_id"`
	Purpose      int     `json:"purpose"`
	Introduction *string `json:"introduction"`
	// Interests    *string `json:"interests"`
	// LookingFor   *string `json:"looking_for"`
	// Basic Profile
	HeightID           int    `json:"height_id"`
	Height             string `json:"height"`
	BodyTypeID         int    `json:"body_type_id"`
	BodyType           string `json:"body_type"`
	BloodTypeID        int    `json:"blood_type_id"`
	BloodType          string `json:"blood_type"`
	ResidenceCountryID int    `json:"residence_country_id"`
	ResidenceCountry   string `json:"residence_country"`
	ResidenceStateID   int    `json:"residence_state_id"`
	ResidenceState     string `json:"residence_state"`
	HometownCountryID  int    `json:"hometown_country_id"`
	HometownCountry    string `json:"hometown_country"`
	HometownStateID    int    `json:"hometown_state_id"`
	HometownState      string `json:"hometown_state"`
	OccupationID       int    `json:"occupation_id"`
	Occupation         string `json:"occupation"`
	EducationID        int    `json:"education_id"`
	Education          string `json:"education"`
	AnnualIncomeID     int    `json:"annual_income_id"`
	AnnualIncome       string `json:"annual_income"`
	SmokingID          int    `json:"smoking_id"`
	Smoking            string `json:"smoking"`
	DrinkingID         int    `json:"drinking_id"`
	Drinking           string `json:"drinking"`
	// User Information
	// Nickname
	// Age
	SiblingsID   int    `json:"siblings_id"`
	Siblings     string `json:"siblings"`
	LanguageID   int    `json:"language_id"`
	Language     string `json:"language"`
	InterestsID  int    `json:"interests_id"`
	Interests    string `json:"interests"`
	LookingForID int    `json:"looking_for_id"`
	LookingFor   string `json:"looking_for"`
	// Education and Career
	SchoolName *string `json:"school_name"`
	JobTitle   *string `json:"job_title"`
	// Attitude towards Love and Marriage
	MaritalHistoryID                 int    `json:"marital_history_id"`
	MaritalHistory                   string `json:"marital_history"`
	PresenceOfChildrenID             int    `json:"presence_of_children_id"`
	PresenceOfChildren               string `json:"presence_of_children"`
	IntentionsTowardsMarriageID      int    `json:"intentions_towards_marriage_id"`
	IntentionsTowardsMarriage        string `json:"intentions_towards_marriage"`
	DesireForChildrenID              int    `json:"desire_for_children_id"`
	DesireForChildren                string `json:"desire_for_children"`
	HouseholdChoresAndChildRearingID int    `json:"household_chores_and_child_rearing_id"`
	HouseholdChoresAndChildRearing   string `json:"household_chores_and_child_rearing"`
	MeetingPreferenceID              int    `json:"meeting_preference_id"`
	MeetingPreference                string `json:"meeting_preference"`
	DatingExpensesID                 int    `json:"dating_expenses_id"`
	DatingExpenses                   string `json:"dating_expenses"`
	// Personality, Lifestyle, Hobbies
	PersonalityTypeID int    `json:"personality_type_id"`
	PersonalityType   string `json:"personality_type"`
	SociabilityID     int    `json:"sociability_id"`
	Sociability       string `json:"sociability"`
	RoommatesID       int    `json:"roommates_id"`
	Roommates         string `json:"roommates"`
	DaysOffID         int    `json:"days_off_id"`
	DaysOff           string `json:"days_off"`
	HobbiesID         int    `json:"hobbies_id"`
	Hobbies           string `json:"hobbies"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
