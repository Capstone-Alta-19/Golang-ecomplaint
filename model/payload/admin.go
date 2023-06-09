package payload

type AddAdminRequest struct {
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=Admin"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginAdminRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginAdminResponse struct {
	Token string `json:"token"`
}

type GetAdminProfileResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UpdateAdminRequest struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	OldPassword     string `json:"old_password" validate:"required,min=6"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
