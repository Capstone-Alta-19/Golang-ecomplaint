package payload

type CreateNews struct {
	NewsName    string `json:"news_name" validate:"required"`
	PhotoURL    string `json:"photo_url" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}

type GetNewsByIDResponse struct {
	ID          uint       `json:"id"`
	NewsName    string     `json:"news_name"`
	PhotoURL    string     `json:"photo_url"`
	Description string     `json:"description"`
	Admin       string     `json:"admin"`
	Category    string     `json:"category"`
	CreatedAt   string     `json:"created_at"`
	NewsList    []NewsList `json:"news_list"`
}

type NewsList struct {
	ID       uint   `json:"id"`
	NewsName string `json:"news_name"`
	PhotoURL string `json:"photo_url"`
}
