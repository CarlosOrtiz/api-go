package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(300);not null" json:"name"`
	Lastname string `gorm:"type:varchar(300);not null" json:"lastname"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Tasks    []Task `json:"tasks"`
}
