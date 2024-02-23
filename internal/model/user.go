package model

import (
	"go-todo/internal/domain"
	"time"
)

type User struct {
	ID        uint      `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) ToDomain() domain.User {
	return domain.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
