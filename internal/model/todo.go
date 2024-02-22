package model

type Todo struct {
}

func (t Todo) TableName() string {
	return "todo"
}
