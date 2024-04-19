package engine

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")
var ErrNoAction = errors.New("no action")
var ErrAlreadyExists = errors.New("already exists")
var ErrRetry = errors.New("retry")

type UpdateFunc func(ctx context.Context, data []byte) ([]byte, error)

type Engine interface {
	// Set only create a new if don't exist.
	Set(ctx context.Context, key string, data []byte) error

	// Get return the store data.
	Get(ctx context.Context, key string) ([]byte, error)

	// Update first get the object and return the value.
	// nil data means that the result don't exist.
	// Returning nil or error will ignore the update.
	// TODO(ramonberrutti): Think a generic locking system.
	Update(ctx context.Context, key string, f UpdateFunc) error

	// Delete remove the resource. ok will return if the resource exist.
	Delete(ctx context.Context, key string) (bool, error)
}
