package todos

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Todo This is a model for TODO
type Todo struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Completed bool      `json:"completed" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Validate validate input data
func (t *Todo) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
