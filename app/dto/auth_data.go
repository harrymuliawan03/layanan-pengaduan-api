package dto

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"roles"`
}

type LoginResponse struct {
	Valid bool   `json:"valid"`
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"roles"`
}
