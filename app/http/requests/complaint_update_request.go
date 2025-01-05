package requests

type ComplaintUpdateCreateRequest struct {
	ComplaintID uint   `json:"complaint_id" form:"complaint_id" validate:"required"`
	AdminID     uint   `json:"admin_id" form:"admin_id" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	Note        string `json:"note" form:"note" validate:"required"`
	UpdatedAt   string `json:"updated_at" form:"updated_at" validate:"required"`
}

type ComplaintUpdateUpdateRequest struct {
	ComplaintID uint   `json:"complaint_id" form:"complaint_id" validate:"required"`
	AdminID     uint   `json:"admin_id" form:"admin_id" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	Note        string `json:"note" form:"note" validate:"required"`
	UpdatedAt   string `json:"updated_at" form:"updated_at" validate:"required"`
}
