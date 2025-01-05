package dto

type ComplaintData struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Photo       string `json:"photo"`
	Address     string `json:"address"`
}
