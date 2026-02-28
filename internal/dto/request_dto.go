package dto

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,numeric,max=20"`
	Password    string `json:"password" binding:"required,min=6"`
}
