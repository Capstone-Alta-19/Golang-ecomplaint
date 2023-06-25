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
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.POST("/register/user", controller.RegisterUserController)
	e.POST("/login/user", controller.LoginUserController)
	e.POST("/login/admin", controller.LoginAdminController)
	e.GET("/news", controller.GetNewsController)
	e.POST("/upload", controller.UploadFileController)

	// user collection
	user := e.Group("/complaintz", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.POST("/complaint", controller.CreateComplaintController)
	user.GET("/complaint", controller.GetUserComplaintsByStatusController)
	user.GET("/complaint/:id", controller.GetComplaintByIDController)
	user.DELETE("/complaint/:id", controller.DeleteComplaintByIDController)
	user.GET("/complaint/category/:id", controller.GetComplaintsByCategoryIDController)
	user.POST("/complaint/:id/comment", controller.CreateCommentByComplaintIDController)

	user.POST("/complaint/:id/pin", controller.PinnedComplaintByIDController)
	user.DELETE("/complaint/:id/pin", controller.UnpinnedComplaintByIDController)
	user.POST("/complaint/:id/like", controller.LikeByComplaintIDController)
	user.DELETE("/complaint/:id/like", controller.UnlikeByComplaintIDController)
	user.GET("/user", controller.GetUserProfileController)
	user.DELETE("/user", controller.DeleteUserController)
	user.GET("/user/complaint/pin", controller.GetPinnedComplaintController)

	user.GET("/news/:id", controller.GetNewsByIDController)
	user.PUT("/user/:id", controller.UpdateUserController)
	user.PUT("/user/password", controller.ChangePasswordController)

	// admin collection
	admin := e.Group("/dashboard", middleware.JWT([]byte(constant.SECRET_JWT)))
	admin.GET("/notification", controller.GetNotificationController)
	admin.POST("/admin", controller.AddAdminController)
	admin.GET("/admin", controller.GetAdminController)
	admin.PUT("/admin", controller.UpdateAdminController)

	admin.POST("/news", controller.CreateNewsController)
	admin.DELETE("/news/:id", controller.DeleteNewsController)
	admin.PUT("/news/:id", controller.UpdateNewsController)

	admin.GET("", controller.GetTotalComplaintsController)
	admin.GET("/complaint", controller.GetAllComplaintsController)
	admin.GET("/complaint/export", controller.ExportComplaintController)
	admin.POST("/complaint/:id", controller.CreateFeedbackByComplaintIDController)
	admin.GET("/complaint/:id", controller.AdminGetComplaintByIDController)
	admin.PUT("/complaint/:id", controller.UpdateComplaintController)
	admin.DELETE("/complaint/:id", controller.AdminDeleteComplaintByIDController)
}
