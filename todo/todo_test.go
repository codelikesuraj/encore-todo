package todo

import (
	"context"
	"testing"
)

func TestSaveTodo(t *testing.T) {
	refreshTodos(t)

	want := "Go to the beach"

	resp, err := SaveTodo(context.Background(), TodoParam{Name: want})
	assertNoError(t, err)

	assertEqual(t, resp.Todo.Name, want)
}

func TestListTodos(t *testing.T) {
	refreshTodos(t)

	ctx := context.Background()
	todos := []TodoParam{
		{Name: "Go to the beach"},
		{Name: "Buy a house"},
		{Name: "Buy a lambo"},
	}

	for _, todo := range todos {
		_, err := SaveTodo(ctx, todo)
		assertNoError(t, err)
	}

	resp, err := ListTodos(ctx)
	assertNoError(t, err)

	got, want := len(resp.Todos), len(todos)
	assertEqual(t, got, want)
}

func TestFetchTodo(t *testing.T) {
	ctx := context.Background()

	want := "Go to the beach now"
	resp, err := SaveTodo(ctx, TodoParam{Name: want})
	assertNoError(t, err)

	assertEqual(t, resp.Todo.Name, want)

	resp, err = FetchTodo(ctx, resp.Todo.ID)
	assertNoError(t, err)

	assertEqual(t, resp.Todo.Name, want)
}

func TestDeleteTodo(t *testing.T) {
	ctx := context.Background()

	want := "Go to the beach now"
	resp, err := SaveTodo(ctx, TodoParam{Name: want})
	assertNoError(t, err)

	resp, err = FetchTodo(ctx, resp.Todo.ID)
	assertNoError(t, err)
	assertEqual(t, resp.Todo.Name, want)

	err = DeleteTodo(ctx, resp.Todo.ID)
	assertNoError(t, err)

	_, err = FetchTodo(ctx, resp.Todo.ID)
	if err == nil {
		t.Fatal("wanted an error")
	}
}

func refreshTodos(t *testing.T) {
	t.Helper()
	Todos = []Todo{}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("wanted %v got %v", want, got)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
