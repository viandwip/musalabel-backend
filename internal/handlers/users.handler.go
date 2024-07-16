package handlers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/oktaviandwip/musalabel/config"
	models "github.com/oktaviandwip/musalabel/internal/models/users"
	"github.com/oktaviandwip/musalabel/internal/repository"
	"github.com/oktaviandwip/musalabel/pkg"
)

type HandlerUsers struct {
	repository.RepoUsersIF
}

func NewUser(r repository.RepoUsersIF) *HandlerUsers {
	return &HandlerUsers{r}
}

// Create User
func (h *HandlerUsers) CreateNewUser(ctx *gin.Context) {
	var err error
	user := models.User{
		Role: "user",
	}

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	result, err := h.CreateUser(&user)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(201, result).Send(ctx)
}

// Update Profile
func (h *HandlerUsers) PostProfile(ctx *gin.Context) {
	var err error
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	user.Image = ctx.MustGet("profileImage").(string)
	result, err := h.UpdateProfile(&user)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}

// // Get Profile for header
// func (h *HandlerUsers) GetProfileForHeader(ctx *gin.Context) {

// 	id := ctx.Param("id")

// 	result, err := h.FetchProfileForHeader(id)
// 	if err != nil {
// 		pkg.NewRes(400, &config.Result{
// 			Data: err.Error(),
// 		}).Send(ctx)
// 		return
// 	}

// 	pkg.NewRes(200, result).Send(ctx)
// }

// // Get Profile
// func (h *HandlerUsers) GetProfile(ctx *gin.Context) {

// 	id := ctx.Param("id")

// 	result, err := h.FetchProfile(id)
// 	if err != nil {
// 		pkg.NewRes(400, &config.Result{
// 			Data: err.Error(),
// 		}).Send(ctx)
// 		return
// 	}

// 	pkg.NewRes(200, result).Send(ctx)
// }
