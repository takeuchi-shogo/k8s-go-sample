package interactor

import (
	"net/http"
	"strconv"
	"time"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/repository"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/services"
)

type UserProfileInteractor struct {
	DBRepository          repository.DBRepository
	UserProfileRepository repository.UserProfileRepository
}

func (interactor *UserProfileInteractor) Get(id int) (*models.UserProfiles, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	profile, err := interactor.UserProfileRepository.FindByID(db, id)
	if err != nil {
		return &models.UserProfiles{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return profile, nil
}

func (interactor *UserProfileInteractor) Create(p *models.UserProfiles) (*models.UserProfiles, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	profile, err := interactor.UserProfileRepository.Create(db, p)
	if err != nil {
		return &models.UserProfiles{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return profile, services.NewResultStatus(http.StatusOK, nil)
}

func (interactor *UserProfileInteractor) Save(up *types.UpdateUserProfiles) (*models.UserProfiles, *services.ResultStatus) {
	db := interactor.DBRepository.Connect()

	id, _ := strconv.Atoi(up.ID)

	foundProfile, err := interactor.UserProfileRepository.FindByID(db, id)
	if err != nil {
		return &models.UserProfiles{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	if up.Introduction != nil {
		if *up.Introduction != "" {
			foundProfile.Introduction = up.Introduction
		}
	}

	if checkValue(up.BodyTypeID) {
		id := *up.BodyTypeID
		foundProfile.BodyTypeID = id
	}
	if checkValue(up.BloodTypeID) {
		id := *up.BloodTypeID
		foundProfile.BloodTypeID = id
	}
	if checkValue(up.ResidenceCountryID) {
		id := *up.ResidenceCountryID
		foundProfile.ResidenceCountryID = id
	}
	if checkValue(up.ResidenceStateID) {
		id := *up.ResidenceStateID
		foundProfile.ResidenceStateID = id
	}
	if checkValue(up.HometownCountryID) {
		id := *up.HometownCountryID
		foundProfile.HometownCountryID = id
	}
	if checkValue(up.HometownStateID) {
		id := *up.HometownStateID
		foundProfile.HometownStateID = id
	}
	if checkValue(up.OccupationID) {
		id := *up.OccupationID
		foundProfile.OccupationID = id
	}
	if checkValue(up.EducationID) {
		id := *up.EducationID
		foundProfile.EducationID = id
	}
	if checkValue(up.AnnualIncomeID) {
		id := *up.AnnualIncomeID
		foundProfile.AnnualIncomeID = id
	}
	if checkValue(up.SmokingID) {
		id := *up.SmokingID
		foundProfile.SmokingID = id
	}
	if checkValue(up.DrinkingID) {
		id := *up.DrinkingID
		foundProfile.DrinkingID = id
	}

	if checkValue(up.SiblingsID) {
		id := *up.SiblingsID
		foundProfile.SiblingsID = id
	}
	if checkValue(up.LanguageID) {
		id := *up.LanguageID
		foundProfile.LanguageID = id
	}
	if checkValue(up.InterestsID) {
		id := *up.InterestsID
		foundProfile.InterestsID = id
	}
	if checkValue(up.LookingForID) {
		id := *up.LookingForID
		foundProfile.LookingForID = id
	}

	if checkValue(up.MaritalHistoryID) {
		id := *up.MaritalHistoryID
		foundProfile.MaritalHistoryID = id
	}
	if checkValue(up.PresenceOfChildrenID) {
		id := *up.PresenceOfChildrenID
		foundProfile.PresenceOfChildrenID = id
	}
	if checkValue(up.IntentionsTowardsMarriageID) {
		id := *up.IntentionsTowardsMarriageID
		foundProfile.IntentionsTowardsMarriageID = id
	}
	if checkValue(up.DesireForChildrenID) {
		id := *up.DesireForChildrenID
		foundProfile.DesireForChildrenID = id
	}
	if checkValue(up.HouseholdChoresAndChildRearingID) {
		id := *up.HouseholdChoresAndChildRearingID
		foundProfile.HouseholdChoresAndChildRearingID = id
	}
	if checkValue(up.MeetingPreferenceID) {
		id := *up.MeetingPreferenceID
		foundProfile.MeetingPreferenceID = id
	}
	if checkValue(up.DatingExpensesID) {
		id := *up.DatingExpensesID
		foundProfile.DatingExpensesID = id
	}

	if checkValue(up.PersonalityTypeID) {
		id := *up.PersonalityTypeID
		foundProfile.PersonalityTypeID = id
	}
	if checkValue(up.SociabilityID) {
		id := *up.SociabilityID
		foundProfile.SociabilityID = id
	}
	if checkValue(up.RoommatesID) {
		id := *up.RoommatesID
		foundProfile.RoommatesID = id
	}
	if checkValue(up.DaysOffID) {
		id := *up.DaysOffID
		foundProfile.DaysOffID = id
	}

	foundProfile.UpdatedAt = time.Now().Unix()

	profile, err := interactor.UserProfileRepository.Save(db, foundProfile)
	if err != nil {
		return &models.UserProfiles{}, services.NewResultStatus(http.StatusBadRequest, err)
	}
	return profile, services.NewResultStatus(http.StatusOK, nil)
}

func checkValue(id *int) bool {
	if id != nil {
		return true
	}
	return false
}
