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
	profile = setUserInformation(profile)
	profile = setAttitudeTowardsLoveAndMarriage(profile)
	profile = setPersonalityAndLifstyleAndHobbies(profile)
	return profile
}

func setBasicProfile(p *models.UserProfiles) *models.ResponseUserProfiles {
	profile := &models.ResponseUserProfiles{
		Introduction:       p.Introduction,
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

		SiblingsID:   p.SiblingsID,
		LanguageID:   p.LanguageID,
		InterestsID:  p.InterestsID,
		LookingForID: p.LookingForID,

		MaritalHistoryID:                 p.MaritalHistoryID,
		PresenceOfChildrenID:             p.PresenceOfChildrenID,
		IntentionsTowardsMarriageID:      p.IntentionsTowardsMarriageID,
		DesireForChildrenID:              p.DesireForChildrenID,
		HouseholdChoresAndChildRearingID: p.HouseholdChoresAndChildRearingID,
		MeetingPreferenceID:              p.MeetingPreferenceID,
		DatingExpensesID:                 p.DatingExpensesID,

		PersonalityTypeID: p.PersonalityTypeID,
		SociabilityID:     p.SociabilityID,
		RoommatesID:       p.RoommatesID,
		DaysOffID:         p.DaysOffID,
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
	p.Siblings = setSiblings(p.SiblingsID)
	p.Interests = setInterests(p.InterestsID)
	p.LookingFor = setLookingFor(p.LookingForID)
	return p
}

func setSiblings(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "長男/長女"
	case 2:
		return "次男/次女"
	case 3:
		return "末っ子"
	case 4:
		return "一人っ子"
	}
	return ""
}

func setInterests(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "メル友/TEL友"
	case 2:
		return "趣味友達"
	case 3:
		return "恋人"
	case 4:
		return "婚約者"
	}
	return ""
}

func setLookingFor(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "メル友"
	case 2:
		return "恋人"
	case 3:
		return "婚活"
	case 4:
		return "その他"
	}
	return ""
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
	p.MeetingPreference = setMeetingPreference(p.MeetingPreferenceID)
	p.DatingExpenses = setDatingExpenses(p.DatingExpensesID)

	return p
}

func setMaritalHistory(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "独身(未婚)"
	case 2:
		return "独身(離婚)"
	case 3:
		return "バツあり"
	case 4:
		return "既婚"
	}

	return ""
}

func setPresenceOfChildren(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "なし"
	case 2:
		return "同居"
	case 3:
		return "別居"
	}
	return ""
}

func setIntentionsTowardsMarriage(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "今すぐしたい"
	case 2:
		return "2〜3年以内にしたい"
	case 3:
		return "いい人がいればしたい"
	case 4:
		return "相手と相談して決めたい"
	case 5:
		return "わからない"
	}
	return ""
}

func setDesireForChildren(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "欲しい"
	case 2:
		return "いらない"
	case 3:
		return "相手と相談して決めたい"
	case 4:
		return "わからない"
	}
	return ""
}

func setHouseholdChoresAndChildRearing(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "積極的に参加したい"
	case 2:
		return "できれば参加したい"
	case 3:
		return "相手と相談して決めたい"
	case 4:
		return "できれば相手に任せたい"
	case 5:
		return "相手に任せたい"
	}
	return ""
}

func setMeetingPreference(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "ますは会いたい"
	case 2:
		return "気が合えば会いたい"
	case 3:
		return "メッセージを重ねてから会いたい"
	case 4:
		return "メールなどやり取りだけ"
	}
	return ""
}

func setDatingExpenses(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "自分が払う"
	case 2:
		return "自分が多く払う"
	case 3:
		return "割り勘"
	case 4:
		return "自分のものは自分で払う"
	case 5:
		return "相手が多く払う"
	case 6:
		return "相手が払う"
	case 7:
		return "相談して決める"
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
		return ""
	case 1:
		return "大人数が好き"
	case 2:
		return "少人数が好き"
	case 3:
		return "一人が好き"
	case 4:
		return "徐々に仲良くなる"
	case 5:
		return "すぐ仲良くなる"
	}
	return ""
}

func setRoommates(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "一人暮らし"
	case 2:
		return "実家暮らし"
	case 3:
		return "友達と暮らしている"
	case 4:
		return "ペットと一緒"
	case 5:
		return "寮、社宅"
	}
	return ""
}

func setDaysOff(id int) string {
	switch id {
	case 0:
		return ""
	case 1:
		return "土日"
	case 2:
		return "平日"
	case 3:
		return "不定休"
	}
	return ""
}
