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

func (up *UserPriflesPresenter) ResponseUserProfile(p *models.UserProfiles) *models.ResponseUserProfiles {
	profile := setBasicProfile(p)
	return profile
}

func setBasicProfile(p *models.UserProfiles) *models.ResponseUserProfiles {
	profile := &models.ResponseUserProfiles{
		HeightID:           p.HeightID,
		BodyTypeID:         p.BodyTypeID,
		BloodTypeID:        p.BloodTypeID,
		ResidenceCountryID: p.ResidenceCountryID,
		ResidenceStateID:   p.ResidenceStateID,
		HometownCountryID:  p.HometownCountryID,
		HometownStateID:    p.HometownStateID,
		OccupationID:       p.OccupationID,
		EducationID:        p.EducationID,
		AnnualIncomeID:     p.AnnualIncomeID,
		SmokingID:          p.SmokingID,
		DrinkingID:         p.DrinkingID,
	}
	profile.Height = setHeight(p.HeightID)
	profile.BodyType = setBodyType(p.BodyTypeID)
	profile.BloodType = setBloodType(p.BloodTypeID)
	profile.ResidenceCountry, profile.ResidenceState = setResidence(p.ResidenceCountryID, p.ResidenceStateID)
	profile.HometownCountry, profile.HometownState = setHometown(p.HometownCountryID, p.HometownStateID)
	profile.Occupation = setOccupation(p.OccupationID)
	profile.Education = setEducation(p.EducationID)
	profile.AnnualIncome = setAnnualIncome(p.AnnualIncomeID)
	profile.Smoking = setSmoking(p.SmokingID)
	profile.Drinking = setDrinking(p.DrinkingID)
	return profile
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
	switch id {
	case 1:
		return "北海道"
	case 2:
		return "青森県"
	case 3:
		return "岩手県"
	case 4:
		return "宮城県"
	case 5:
		return "秋田県"
	case 6:
		return "山形県"
	case 7:
		return "福島県"
	case 8:
		return "茨城県"
	case 9:
		return "栃木県"
	case 10:
		return "群馬県"
	case 11:
		return "埼玉県"
	case 12:
		return "千葉県"
	case 13:
		return "東京都"
	case 14:
		return "神奈川県"
	case 15:
		return "新潟県"
	case 16:
		return "富山県"
	case 17:
		return "石川県"
	case 18:
		return "福井県"
	case 19:
		return "山梨県"
	case 20:
		return "長野県"
	case 21:
		return "岐阜県"
	case 22:
		return "静岡県"
	}
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
