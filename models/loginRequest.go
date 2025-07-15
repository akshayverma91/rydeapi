package models

// swagger:model LoginRequest
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
