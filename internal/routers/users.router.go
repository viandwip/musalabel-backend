package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/oktaviandwip/musalabel/internal/handlers"
	"github.com/oktaviandwip/musalabel/internal/middleware"
	"github.com/oktaviandwip/musalabel/internal/repository"
)

func users(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/users")

	repo := repository.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", handler.CreateNewUser)
	route.PATCH("/profile/", middleware.Authjwt("admin", "user"), middleware.UploadFile, handler.PostProfile)
	// route.GET("/profile/:id", middleware.Authjwt("admin", "user"), handler.GetProfile)
	// route.GET("/profile/header/:id", middleware.Authjwt("user"), handler.GetProfileForHeader)

}
