package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oktaviandwip/musalabel/config"
	models "github.com/oktaviandwip/musalabel/internal/models/users"
	"github.com/oktaviandwip/musalabel/internal/repository"
	"github.com/oktaviandwip/musalabel/pkg"
)

type HandlerAuth struct {
	*repository.RepoUsers
}

func NewAuth(r *repository.RepoUsers) *HandlerAuth {
	return &HandlerAuth{r}
}

// Login
func (h *HandlerAuth) Login(ctx *gin.Context) {
	var data models.User

	if err := ctx.ShouldBind(&data); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	user, err := h.GetPassByEmail(data.Email)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(user.Password, data.Password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "Password incorrect",
		}).Send(ctx)
		return
	}

	jwt := pkg.NewToken(user.Id, user.Role)
	token, err := jwt.Generate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	type result struct {
		Token string
		User  *models.User
	}

	response := result{
		Token: token,
		User:  user,
	}

	pkg.NewRes(200, &config.Result{Data: response}).Send(ctx)
}
