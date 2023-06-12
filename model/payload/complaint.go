package payload

import "time"

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

type GetComplaintByCategoryIDResponse struct {
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

type GetTotalComplaintsResponse struct {
	Total      uint `json:"total"`
	Complaint  uint `json:"complaint"`
	Aspiration uint `json:"aspiration"`
}
