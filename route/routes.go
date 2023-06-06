package route

import (
	"capstone/constant"
	"capstone/controller"
	"capstone/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.POST("/register/user", controller.RegisterUserController)
	e.POST("/login/user", controller.LoginUserController)
	e.POST("/login/admin", controller.LoginAdminController)
	e.GET("/news", controller.GetNewsController)

	// user collection
	user := e.Group("/user", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.GET("complaint/:id", controller.GetComplaintController)
	user.GET("feedback/:complaintID", controller.GetFeedbackController)
	user.POST("/complaint", controller.CreateComplaintController)
	user.GET("/news/:id", controller.GetNewsController)
	user.PUT("/username", controller.UpdateUserController)
	user.PUT("/password", controller.UpdateUserController)
	user.PUT("/name", controller.UpdateUserController)
	user.PUT("/photoprofile", controller.UpdateUserController)
	user.PUT("/phone", controller.UpdateUserController)
	user.PUT("/email", controller.UpdateUserController)

	// admin collection
	admin := e.Group("/admin", middleware.JWT([]byte(constant.SECRET_JWT)))
	admin.POST("/add", controller.AddAdminController)
	admin.POST("/news", controller.CreateNewsController)
	admin.GET("/news/:id", controller.GetNewsController)
	admin.DELETE("/news", controller.DeleteNewsController)
	admin.PUT("/news", controller.UpdateNewsController)

	admin.GET("/complaint/:id", controller.GetComplaintByIDController)
}
