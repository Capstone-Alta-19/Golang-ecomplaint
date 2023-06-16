package payload

type CreateNews struct {
	NewsName    string `json:"news_name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}
