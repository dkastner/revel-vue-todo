package controllers

import (
	"fmt"
	"todo/app/entities"
	"todo/app/repositories"

	"github.com/revel/revel"
)

type TodosController struct {
	ApplicationController
	todos *repositories.TodosRepository
}

type todoItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c TodosController) Index() revel.Result {
	repo := repositories.TodosRepository{DB: c.DB}

	return c.RenderJSON(jsonifyTodos(repo.AllTodos()))
}

func (c TodosController) Create() revel.Result {
	c.Response.Status = 201

	repo := repositories.TodosRepository{DB: c.DB}

	fmt.Println(string(c.Params.JSON))

	var newTodo entities.TodoJson
	c.Params.BindJSON(&newTodo)

	fmt.Println(newTodo)

	repo.CreateTodo(newTodo.Name)

	return c.RenderJSON(jsonifyTodos(repo.AllTodos()))
}

func jsonifyTodos(todos []entities.Todo) []entities.TodoJson {
	var jsonTodos []entities.TodoJson
	for _, todo := range todos {
		jsonTodos = append(jsonTodos, entities.TodoJson{ID: todo.ID, Name: todo.Name})
	}

	return jsonTodos
}
