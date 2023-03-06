package presenters

import (
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/presenter"
)

type UserPriflesPresenter struct {
}

func NewUserProfilesPresenter() presenter.UserProfilePresenter {
	return &UserPriflesPresenter{}
}

func (up *UserPriflesPresenter) ResponseUserProfile(profile *models.ResponseUserProfiles) *models.ResponseUserProfiles {
	profile = setBasicProfile(profile)
	return profile
}

func setBasicProfile(p *models.ResponseUserProfiles) *models.ResponseUserProfiles {
	p.Height = setHeight(p.HeightID)
	p.BodyType = setBodyType(p.BodyTypeID)
	p.BloodType = setBloodType(p.BloodTypeID)
	p.ResidenceCountry, p.ResidenceState = setResidence(p.ResidenceCountryID, p.ResidenceStateID)
	p.HometownCountry, p.HometownState = setHometown(p.HometownCountryID, p.HometownStateID)
	p.Occupation = setOccupation(p.OccupationID)
	p.Education = setEducation(p.EducationID)
	p.AnnualIncome = setAnnualIncome(p.AnnualIncomeID)
	p.Smoking = setSmoking(p.SmokingID)
	p.Drinking = setDrinking(p.DrinkingID)

	return p
}

func setHeight(id int) string {
	switch id {
	case 0:
		return "こだわらない"
	case 1:
		return "170cm以下"
	case 2:
		return "170cm以上"
	}
	return ""
}

func setBodyType(id int) string {
	switch id {
	case 0:
		return "指定なし"
	case 1:
		return "細い"
	case 2:
		return "普通"
	case 3:
		return "太い"
	}
	return ""
}

func setBloodType(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "A型"
	case 2:
		return "B型"
	case 3:
		return "O型"
	case 4:
		return "AB型"
	case 5:
		return "こだわらない"
	}
	return ""
}

func setResidence(countryID, stateID int) (string, string) {
	country, state := "", ""
	country = setCountry(countryID)
	state = setState(countryID, stateID)
	return country, state
}

func setHometown(countryID, stateID int) (string, string) {
	country, state := "", ""
	country = setCountry(countryID)
	state = setState(countryID, stateID)
	return country, state
}

func setCountry(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "日本"
	case 2:
		return "海外"
	}
	return ""
}

func setState(countryID, stateID int) string {
	switch countryID {
	case 1:
		// 国内の県名
		return setPrefecture(stateID)
	case 2:
		// 海外の国名
		return setCountryName(stateID)
	}
	return ""
}

func setPrefecture(id int) string {
	return ""
}

func setCountryName(id int) string {
	return ""
}

func setOccupation(id int) string {
	return ""
}

func setEducation(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "中卒"
	case 2:
		return "高卒"
	case 3:
		return "高専卒"
	case 4:
		return "大卒"
	case 5:
		return "大学院卒"
	}
	return ""
}

func setAnnualIncome(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "200万〜400万"
	case 2:
		return "400万〜600万"
	case 3:
		return "600万〜800万"
	case 4:
		return "800万以上"
	}
	return ""
}

func setSmoking(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "吸わない"
	case 2:
		return "吸う"
	}
	return ""
}

func setDrinking(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "飲まない"
	case 2:
		return "たまに飲む"
	case 3:
		return "飲む"
	}
	return ""
}

func setUserInformation(p *models.ResponseUserProfiles) *models.ResponseUserProfiles {
	return p
}

func setEducationAndCareer(p *models.ResponseUserProfiles) *models.ResponseUserProfiles {
	return p
}

func setAttitudeTowardsLoveAndMarriage(p *models.ResponseUserProfiles) *models.ResponseUserProfiles {
	p.MaritalHistory = setMaritalHistory(p.MaritalHistoryID)
	p.PresenceOfChildren = setPresenceOfChildren(p.PresenceOfChildrenID)
	p.IntentionsTowardsMarriage = setIntentionsTowardsMarriage(p.IntentionsTowardsMarriageID)
	p.DesireForChildren = setDesireForChildren(p.DesireForChildrenID)
	p.HouseholdChoresAndChildRearing = setHouseholdChoresAndChildRearing(p.HouseholdChoresAndChildRearingID)
	p.IdealFirstEncounter = setIdealFirstEncounter(p.IdealFirstEncounterID)
	p.DatingExpenses = setDatingExpenses(p.DatingExpensesID)

	return p
}

func setMaritalHistory(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setPresenceOfChildren(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setIntentionsTowardsMarriage(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setDesireForChildren(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setHouseholdChoresAndChildRearing(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setIdealFirstEncounter(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setDatingExpenses(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setPersonalityAndLifstyleAndHobbies(p *models.ResponseUserProfiles) *models.ResponseUserProfiles {
	p.PersonalityType = setPersonalityType(p.PersonalityTypeID)
	p.Sociability = setSociability(p.SociabilityID)
	p.Roommates = setRoommates(p.RoommatesID)
	p.DaysOff = setDaysOff(p.DaysOffID)

	return p
}

func setPersonalityType(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setSociability(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setRoommates(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}

func setDaysOff(id int) string {
	switch id {
	case 0:
	case 1:
	}
	return ""
}
