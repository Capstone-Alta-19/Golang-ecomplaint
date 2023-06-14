package payload

type CreateFeedbackRequest struct {
	Description string `json:"description" validate:"required"`
}
