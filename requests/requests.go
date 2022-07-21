package requests

type LoginRequest struct {
	Email    string `json:"email" validation:"required"`
	Password string `json:"password" validation:"required"`
}
