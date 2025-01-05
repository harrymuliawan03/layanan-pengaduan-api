package requests

type ComplaintCreateRequest struct {
	UserID      uint   `json:"user_id" form:"user_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	Photo       string `json:"photo" form:"photo" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
}

type ComplaintUpdateRequest struct {
	UserID      uint   `json:"user_id" form:"user_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	Photo       string `json:"photo" form:"photo" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
}
