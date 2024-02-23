package model

import (
	"go-todo/internal/domain"
	"time"
)

type Todo struct {
	ID          uint      `gorm:"column:id"`
	CreatorID   uint      `gorm:"column:creator_id"`
	CategoryID  uint      `gorm:"column:category_id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Priority    string    `gorm:"column:priority"`
	Complete    bool      `gorm:"column:complete"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	User        User      `gorm:"foreignKey:creator_id"`
	Category    Category  `gorm:"foreignKey:category_id"`
}

func (t Todo) TableName() string {
	return "todo"
}

func (t Todo) ToDomain() domain.Todo {
	return domain.Todo{
		ID:          t.ID,
		CreatorID:   t.CreatorID,
		CategoryID:  t.CategoryID,
		Title:       t.Title,
		Description: t.Description,
		Priority:    domain.TodoPriority(t.Priority),
		Complete:    t.Complete,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
