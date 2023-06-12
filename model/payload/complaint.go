package payload

type CreateComplaintRequest struct {
	Type        string `json:"type" validate:"required,oneof=Complaint Aspiration"`
	Description string `json:"description" validate:"required"`
	PhotoURL    string `json:"photo_url"`
	VideoURL    string `json:"video_url"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	IsPublic    bool   `json:"is_public" validate:"omitempty"`
}

type GetComplaintByStatusResponse struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
