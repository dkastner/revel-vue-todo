package entities

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

type TodoJson struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
