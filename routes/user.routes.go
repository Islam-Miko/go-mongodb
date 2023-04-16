package routes

import (
	"github.com/Islam-Miko/go-mongodb/controllers"
	"github.com/Islam-Miko/go-mongodb/services"
	"github.com/Islam-Miko/go-mongodb/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController  controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("/users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}