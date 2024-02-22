package model

type Category struct {
}

func (c Category) TableName() string {
	return "category"
}
