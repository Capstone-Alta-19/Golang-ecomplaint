package payload

type GetComplaintRequest struct {
	Description string `json:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	Feedback    string `json:"feedback" validate:"required"`
}
