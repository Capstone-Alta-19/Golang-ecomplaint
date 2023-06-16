package payload

type CreateCommentRequest struct {
	Description string `json:"description" validate:"required"`
}

type GetCommentResponse struct {
	ID           uint    `json:"id"`
	PhotoProfile *string `json:"photo_profile"`
	FullName     string  `json:"full_name"`
	Username     string  `json:"username"`
	Description  string  `json:"description"`
	CreatedAt    string  `json:"created_at"`
}
