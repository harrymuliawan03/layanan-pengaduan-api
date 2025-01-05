package dto

type ComplaintUpdateData struct {
	ID          uint   `json:"id"`
	ComplaintID uint   `json:"complaint_id"`
	AdminID     uint   `json:"admin_id"`
	Status      string `json:"status"`
	Note        string `json:"note"`
	UpdatedAt   string `json:"updated_at"`
}
