package requests

type RegisterRequest struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	// Phone                string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
	// PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Username  string `json:"username" validate:"required"`
	FirstName string `json:"first_name" form:"first_name" validate:"required"`
	LastName  string `json:"last_name" form:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type VerifEmailRequest struct {
	Verified bool `json:"verified" form:"verified" validate:"required,boolean"`
}
