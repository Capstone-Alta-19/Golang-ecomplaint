package payload

import "time"

type CreateUserRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type LoginUserRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type OtpEmailRequest struct {
	Email string `json:"email"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type GetUserProfileResponse struct {
	ID           uint    `json:"id"`
	PhotoProfile *string `json:"photo_profile"`
	FullName     string  `json:"full_name"`
	Laporan      uint    `json:"laporan"`
	Pending      uint    `json:"pending"`
	Proccess     uint    `json:"proccess"`
	Resolved     uint    `json:"resolved"`
}

type GetUserComplaintIDResponse struct {
	ID           uint      `json:"id"`
	PhotoProfile *string   `json:"photo_profile"`
	FullName     string    `json:"full_name"`
	Username     string    `json:"username"`
	Category     string    `json:"category"`
	Description  string    `json:"description"`
	PhotoURL     *string   `json:"photo_url"`
	VideoURL     *string   `json:"video_url"`
	IsPublic     bool      `json:"is_public"`
	Feedback     *string   `json:"feedback"`
	LikesCount   uint      `json:"likes_count"`
	CreatedAt    time.Time `json:"created_at"`
}
