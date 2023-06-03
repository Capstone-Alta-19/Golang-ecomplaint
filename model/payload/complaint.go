package payload

type CreateComplaintRequest struct {
	Type        string `json:"type" validate:"required,oneof=Complaint Aspiration"`
	Description string `json:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	IsPublic    bool   `json:"is_public" validate:"required"`
}

type GetComplaintRequest struct {
	Description string `json:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	Feedback    string `json:"feedback" validate:"required"`
}
