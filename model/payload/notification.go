package payload

type GetNotificationResponse struct {
	ID           uint    `json:"id"`
	UserID       uint    `json:"user_id"`
	PhotoProfile *string `json:"photo_profile"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
	ComplaintID  uint    `json:"complaint_id"`
	CreatedAt    string  `json:"created_at"`
}
