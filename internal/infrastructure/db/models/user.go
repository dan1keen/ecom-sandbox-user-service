package models

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Phone     string    `gorm:"column:phone;size:20;unique" json:"phone"`
	Email     *string   `gorm:"column:email;unique" json:"email"`
	FirstName *string   `gorm:"column:firstname;size:255" json:"firstname"`
	LastName  *string   `gorm:"column:lastname;size:255" json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
