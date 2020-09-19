package todos

import (
	"github.com/gofiber/fiber/v2"
)

var todos = []Todo{
	{ID: 1, Name: "Food to dog", Completed: false},
	{ID: 2, Name: "Food to cat", Completed: false},
}

// GetTodos Fetch all Todos
func GetTodos(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(todos)
}
