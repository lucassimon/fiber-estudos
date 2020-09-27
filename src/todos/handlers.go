package todos

import (
	"github.com/gofiber/fiber/v2"
)

type GetTodosPatameters struct {
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

// GetTodos Fetch all Todos
func GetTodos(context *fiber.Ctx) error {
	parameters := new(GetTodosPatameters)

	if err := context.QueryParser(parameters); err != nil {
		return context.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	todos, err := All(parameters.Limit, parameters.Offset)

	if err != nil {
		return context.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	return context.Status(fiber.StatusOK).JSON(todos)
}

// CreateTodo POST a new TODO item
func CreateTodo(context *fiber.Ctx) error {
	// New TODO struct
	todoPayload := new(Todo)

	// Parse body into struct
	if err := context.BodyParser(todoPayload); err != nil {
		return context.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	todoInserted, err := Create(*todoPayload)

	if err != nil {
		return context.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	return context.Status(fiber.StatusOK).JSON(todoInserted)
}
