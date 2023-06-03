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
	user.GET("complaint/id", controller.GetComplaintController)
	user.POST("/complaint", controller.CreateComplaintController)
	user.GET("/news", controller.GetNewsController)

	// admin collection
	admin := e.Group("/admin", middleware.JWT([]byte(constant.SECRET_JWT)))
	admin.POST("/add", controller.AddAdminController)
	admin.POST("/news", controller.CreateNewsController)
	admin.DELETE("/news", controller.DeleteNewsController)
	admin.PUT("/news", controller.UpdateNewsController)

	// complaint
	e.GET("/complaint/:id", controller.GetComplaintController)

	// comment
	e.GET("/comment", controller.GetCommentController)
}
