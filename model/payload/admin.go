package payload

type AddAdminRequest struct {
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginAdminRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=5"`
}

type LoginAdminResponse struct {
	Token string `json:"token"`
}
