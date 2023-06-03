package payload

type CreateCommentRequest struct {
	Description string `json:"description" validate:"required"`
	NewsID      string `json:"news_id" validate:"required"`
}

type GetCommentRequest struct {
	Description string `json:"description" validate:"required"`
	NewsID      string `json:"news_id" validate:"required"`
}
