package tests

import (
	"encoding/json"
	"strings"
	"todo/app/entities"
	"todo/app/repositories"

	"github.com/revel/modules/orm/gorm/app"

	"github.com/revel/revel/testing"
)

type TodosTest struct {
	testing.TestSuite
	todos *repositories.TodosRepository
}

func setup(t *TodosTest) {
	t.todos = &repositories.TodosRepository{DB: gormdb.DB}

	t.todos.Clear()
}

func teardown(t *TodosTest) {
	t.todos.Clear()
}

type todosResponseJsonItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (t *TodosTest) TestGetTodos() {
	setup(t)

	t1 := t.todos.CreateTodo("Do the dumb things I gotta do")
	t2 := t.todos.CreateTodo("Touch the puppet head")

	t.Get("/api/todos")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	expected := []entities.TodoJson{
		entities.TodoJson{
			ID:   t1.ID,
			Name: t1.Name,
		},
		entities.TodoJson{
			ID:   t2.ID,
			Name: t2.Name,
		},
	}

	actual := []entities.TodoJson{}
	if err := json.Unmarshal(t.ResponseBody, &actual); err != nil {
		panic(err)
	}

	t.AssertEqual(expected, actual)

	teardown(t)
}

func (t *TodosTest) TestPostTodos() {
	setup(t)

	name := "Put your hand inside the puppet head"
	body := `{"name": "` + name + `"}`

	t.Post("/api/todos", "application/json", strings.NewReader(body))
	t.AssertStatus(201)
	t.AssertContentType("application/json; charset=utf-8")

	responseTodos := []entities.TodoJson{}
	if err := json.Unmarshal(t.ResponseBody, &responseTodos); err != nil {
		panic(err)
	}

	t.AssertEqual(1, len(responseTodos))
	t.AssertEqual(name, responseTodos[0].Name)

	records := t.todos.AllTodos()

	t.AssertEqual(1, len(records))
	t.AssertEqual(name, records[0].Name)

	teardown(t)
}
