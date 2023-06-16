package payload

type GetPinnedComplaintResponse struct {
	ID           uint    `json:"id"`
	PhotoProfile *string `json:"photo_profile"`
	FullName     string  `json:"full_name"`
	Username     string  `json:"username"`
	Description  string  `json:"description"`
	PhotoURL     *string `json:"photo_url"`
	VideoURL     *string `json:"video_url"`
	Feedback     *string `json:"feedback"`
	LikesCount   uint    `json:"likes_count"`
	CreatedAt    string  `json:"created_at"`
}
