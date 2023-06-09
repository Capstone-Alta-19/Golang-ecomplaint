package payload

type CreateComplaintRequest struct {
	Type        string `json:"type" validate:"required,oneof=Complaint Aspiration"`
	Description string `json:"description" validate:"required,max=150"`
	PhotoURL    string `json:"photo_url"`
	VideoURL    string `json:"video_url"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	IsPublic    *bool  `json:"is_public" validate:"required"`
}

type GetComplaintByStatusResponse struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type GetComplaintByCategoryIDResponse struct {
	ID           uint    `json:"id"`
	PhotoProfile *string `json:"photo_profile"`
	FullName     string  `json:"full_name"`
	Username     string  `json:"username"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
	PhotoURL     *string `json:"photo_url"`
	VideoURL     *string `json:"video_url"`
	IsPublic     bool    `json:"is_public"`
	Feedback     *string `json:"feedback"`
	LikesCount   uint    `json:"likes_count"`
	CreatedAt    string  `json:"created_at"`
}

type GetTotalComplaintsResponse struct {
	Total      uint `json:"total"`
	Complaint  uint `json:"complaint"`
	Aspiration uint `json:"aspiration"`
}

type GetAllComplaintsResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Status      string `json:"status"`
	IsPublic    bool   `json:"is_public"`
	CreatedAt   string `json:"created_at"`
}

type GetComplaintByIDResponse struct {
	ID          uint    `json:"id"`
	FullName    string  `json:"full_name"`
	Type        string  `json:"type"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	PhotoURL    *string `json:"photo_url"`
	VideoURL    *string `json:"video_url"`
	IsPublic    bool    `json:"is_public"`
	CreatedAt   string  `json:"created_at"`
}

type UpdateComplaintRequest struct {
	Status string `json:"status" validate:"required,oneof=Pending Proccess Resolved"`
	Type   string `json:"type" validate:"required,oneof=Complaint Aspiration"`
}

type GetUserComplaintIDResponse struct {
	ID           uint                 `json:"id"`
	PhotoProfile *string              `json:"photo_profile"`
	FullName     string               `json:"full_name"`
	Username     string               `json:"username"`
	Description  string               `json:"description"`
	PhotoURL     *string              `json:"photo_url"`
	VideoURL     *string              `json:"video_url"`
	IsPublic     bool                 `json:"is_public"`
	Feedback     *string              `json:"feedback"`
	CreatedAt    string               `json:"created_at"`
	Comments     []GetCommentResponse `json:"comments"`
	UserProfile  *string              `json:"user_profile"`
}
