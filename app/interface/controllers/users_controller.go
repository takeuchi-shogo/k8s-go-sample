package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
// 	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways"
// 	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
// 	"github.com/takeuchi-shogo/k8s-go-sample/interface/helpers"
// 	"github.com/takeuchi-shogo/k8s-go-sample/interface/presenters"
// 	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
// )

// type usersController struct {
// 	AuthorizeInteractor interactor.AuthorizeInteractor
// 	Interactor          interactor.UserInteractor
// }

// type UsersController interface {
// 	Get()
// 	Post()
// }

// type UsersControllerProvider struct {
// 	DB  repositories.DB
// 	Jwt gateways.Jwt
// }

// func NewUsersController(p UsersControllerProvider) *usersController {
// 	return &usersController{
// 		AuthorizeInteractor: interactor.AuthorizeInteractor{
// 			Jwt: &gateways.JwtGateway{Jwt: p.Jwt},
// 		},
// 		Interactor: interactor.UserInteractor{
// 			DBRepository:   &repositories.DBRepository{DB: p.DB},
// 			UserRepository: &repositories.UserRepository{},
// 			UserPresenter:  &presenters.UsersPresenter{},
// 		},
// 	}
// }

// func (controller *usersController) Get(c helpers.Context) {
// 	// id, err := strconv.Atoi(c.Param("id"))
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, helpers.NewResponseError(http.StatusBadRequest, err, err.Error()))
// 	// 	return
// 	// }

// 	// user, res := controller.Interactor.Get(id)
// 	// if res.Error != nil {
// 	// 	c.JSON(res.Code, helpers.NewResponseError(res.Code, res.Error, res.Error.Error()))
// 	// 	return
// 	// }
// 	// c.JSON(res.Code, helpers.NewResponseSuccess("success", user))
// }

// func (controller *usersController) Post(c helpers.Context) {
// 	accountID, err := strconv.Atoi(c.PostForm("account_id"))
// 	if err != nil {
// 		code := http.StatusBadRequest
// 		c.JSON(code, helpers.NewResponseError(code, err, err.Error()))
// 		return
// 	}

// 	code := c.PostForm("code")
// 	if err != nil {
// 		code := http.StatusBadRequest
// 		c.JSON(code, helpers.NewResponseError(code, err, err.Error()))
// 		return
// 	}

// 	u := &models.Users{}

// 	if err := c.BindJSON(u); err != nil {
// 		code := http.StatusBadRequest
// 		c.JSON(code, helpers.NewResponseError(code, err, err.Error()))
// 		return
// 	}

// 	user, res := controller.Interactor.Create(u, accountID, code)
// 	if res.Error != nil {
// 		c.JSON(res.Code, helpers.NewResponseError(res.Code, res.Error, res.Error.Error()))
// 		return
// 	}
// 	c.JSON(res.Code, helpers.NewResponseSuccess("success", user))
// }

// func (controller *usersController) Patch(c helpers.Context) {
// 	_, res := controller.AuthorizeInteractor.Verify(c.GetHeader("Authorizetion"))
// 	if res.Error != nil {
// 		c.JSON(res.Code, helpers.NewResponseError(res.Code, res.Error, res.Error.Error()))
// 		return
// 	}

// 	u := &models.Users{}

// 	if err := c.BindJSON(u); err != nil {
// 		code := http.StatusBadRequest
// 		c.JSON(code, helpers.NewResponseError(code, err, err.Error()))
// 		return
// 	}

// 	user, res := controller.Interactor.Save(u)
// 	if res.Error != nil {
// 		c.JSON(res.Code, helpers.NewResponseError(res.Code, res.Error, res.Error.Error()))
// 		return
// 	}
// 	c.JSON(res.Code, helpers.NewResponseSuccess("success", user))
// }
