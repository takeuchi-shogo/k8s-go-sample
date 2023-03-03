package models

type UserProfiles struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	Introduction *string `json:"introduction"`
	Interests    *string `json:"interests"`
	LookingFor   *string `json:"looking_for"`
	// Basic Profile
	HeightID       int `json:"height_id"`
	BodyTypeID     int `json:"body_type_id"`
	BloodTypeID    int `json:"blood_type_id"`
	ResidenceID    int `json:"residence_id"`
	HometownID     int `json:"hometown_id"`
	OccupationID   int `json:"occupation_id"`
	EducationID    int `json:"education_id"`
	AnnualIncomeID int `json:"annual_income_id"`
	SmokingID      int `json:"smoking_id"`
	DrinkingID     int `json:"drinking_id"`
	// User Information
	// Nickname
	// Age
	SiblingsID   int `json:"siblings_id"`
	LanguageID   int `json:"language_id"`
	InterestsID  int `json:"interests_id"`
	LookingForID int `json:"looking_for_id"`
	// Education and Career
	EducationalID int `json:"educational_id"`
	JobTitleID    int `json:"job_title_id"`
	// Attitude towards Love and Marriage
	MaritalHistoryID                 int `json:"marital_history_id"`
	PresenceOfChildrenID             int `json:"presence_of_children_id"`
	IntentionsTowardsMarriageID      int `json:"intentions_towards_marriage_id"`
	DesireForChildrenID              int `json:"desire_for_children_id"`
	HouseholdChoresAndChildRearingID int `json:"household_chores_and_child_rearing_id"`
	IdealFirstEncounterID            int `json:"indeal_first_encointer_id"`
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
