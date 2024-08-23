package todo

import (
	"context"

	"encore.dev/beta/errs"
)

type Todo struct {
	ID        uint
	Name      string
	Completed bool
}

type TodosResp struct{ Todos []Todo }

type TodoResp struct{ Todo Todo }

type TodoParam struct{ Name string }

var (
	Id    uint
	Todos []Todo
)

// encore:api public method=GET path=/todos
func ListTodos(ctx context.Context) (*TodosResp, error) {
	return &TodosResp{Todos: Todos}, nil
}

// encore:api public method=POST path=/todos
func SaveTodo(ctx context.Context, todo TodoParam) (*TodoResp, error) {
	Id += 1
	t := Todo{
		ID:   Id,
		Name: todo.Name,
	}
	Todos = append(Todos, t)
	return &TodoResp{Todo: t}, nil
}

// encore:api public method=GET path=/todos/:id
func FetchTodo(ctx context.Context, id uint) (*TodoResp, error) {
	for _, t := range Todos {
		if t.ID == id {
			return &TodoResp{Todo: t}, nil
		}
	}

	return nil, &errs.Error{
		Code:    errs.NotFound,
		Message: "todo not found",
	}
}

// encore:api public method=DELETE path=/todos/:id
func DeleteTodo(ctx context.Context, id uint) error {
	for i, t := range Todos {
		if t.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return nil
		}
	}

	return &errs.Error{
		Code:    errs.NotFound,
		Message: "todo not found",
	}
}
