package model

type UserDTO struct {
	ID       uint   `json:"userid"`
	Fullname string `json:"full_name"`
	Email    string `json:"email_address"`
}
type UserCreatingDTO struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Language string `json:"lang"`
}
