package dto

type UserDTO struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	LastName string `json:"lastname,omitempty" validate:"omitempty,min=2,max=100"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
}
