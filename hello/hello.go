package hello

import (
	"context"
)

type Response struct {
	Message string
}

// encore:api method=POST path=/hello/:name
func World(ctx context.Context, name string) (*Response, error) {
	msg := "Hello " + name + "!"
	return &Response{Message: msg}, nil
}
