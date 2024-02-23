package model

import (
	"go-todo/internal/domain"
	"time"
)

type Category struct {
	ID        uint      `gorm:"column:id"`
	CreatorID string    `gorm:"column:creator_id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      User      `gorm:"foreignKey:creator_id"`
}

func (c Category) TableName() string {
	return "category"
}

func (c Category) ToDomain() domain.Category {
	return domain.Category{
		ID:        c.ID,
		CreatorID: c.CreatorID,
		Name:      c.Name,
	}
}
