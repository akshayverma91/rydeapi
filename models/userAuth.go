package models

// swagger:model AuthUser
type AuthUser struct {
	User     User   `json:"user" bson:"user"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}
