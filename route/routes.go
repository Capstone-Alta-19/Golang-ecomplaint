package route

import (
	"capstone/constant"
	"capstone/controller"
	"capstone/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.POST("/register/user", controller.RegisterUserController)
	e.POST("/login/user", controller.LoginUserController)
	e.POST("/login/admin", controller.LoginAdminController)

	// user collection
	user := e.Group("/user", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.GET("complaint/id", controller.GetUsersController)
	user.POST("/complaint", controller.CreateComplaintController)

	// admin collection
	admin := e.Group("/admin", middleware.JWT([]byte(constant.SECRET_JWT)))
	admin.GET("/users", controller.GetUsersController)
	admin.POST("", controller.AddAdminController)

}
