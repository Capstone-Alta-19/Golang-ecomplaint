package payload

type CreateCommentRequest struct {
	Description string `json:"description" validate:"required"`
}
