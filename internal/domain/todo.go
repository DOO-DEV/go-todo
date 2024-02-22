package domain

import "time"

type TodoPriority string

const (
	TodoPriorityHigh   TodoPriority = "HIGH"
	TodoPriorityMedium TodoPriority = "MEDIUM"
	TodoPriorityLow    TodoPriority = "LOW"
)

type Todo struct {
	ID          uint
	CreatorID   uint
	CategoryID  uint
	Title       string
	Description string
	Priority    TodoPriority
	Complete    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
