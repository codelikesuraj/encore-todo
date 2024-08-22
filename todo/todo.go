package todo

import "context"

type Todo struct {
	Name string
	Done bool
}

type TodoIndexResp struct {
	Todos []Todo
}

var Todos []Todo

// encore:api public method=GET path=/todos
func Index(ctx context.Context) (*TodoIndexResp, error) {
	return &TodoIndexResp{Todos: Todos}, nil
}
