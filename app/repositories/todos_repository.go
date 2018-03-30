package repositories

import (
	"todo/app/entities"

	gorm "github.com/jinzhu/gorm"
)

type TodosRepository struct {
	DB *gorm.DB
}

func (r TodosRepository) AllTodos() []entities.Todo {
	var todos []entities.Todo

	r.DB.Find(&todos)

	return todos
}

func (r TodosRepository) CreateTodo(name string) entities.Todo {
	todo := entities.Todo{Name: name}
	r.DB.Create(&todo)
	return todo
}

func (r TodosRepository) Clear() {
	r.DB.Unscoped().Delete(&entities.Todo{})
}
