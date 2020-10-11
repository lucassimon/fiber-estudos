package todos

import (
	"github.com/gofiber/fiber/v2"
)

// TodosRouterV1 endpoints
func TodosRouterV1(apiVersion fiber.Router) {
	todoRepo := NewRepo()
	todoService := NewService(todoRepo)

	apiVersion.Get("/todos", FetchTodos(todoService))
	apiVersion.Post("/todos", CreateTodo(todoService))
	apiVersion.Get("/todos/:id", Get(todoService))
	// apiVersion.Put("/todos/:id", UpdateTodo(todoService))
	// apiVersion.Delete("/todos/:id", RemoveTodo(todoService))
}

// GetTodosPatameters paremeters to manipulate query parameters
type GetTodosPatameters struct {
	Limit   int64  `query:"limit"`
	Offset  int64  `query:"offset"`
	OrderBy string `query:"orderby"`
	Asc     string `query:"asc"`
}

// FetchTodos Fetch all Todos
func FetchTodos(service Service) fiber.Handler {
	return func(context *fiber.Ctx) error {
		parameters := new(GetTodosPatameters)

		if err := context.QueryParser(parameters); err != nil {
			return context.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		todos, err := service.FetchTodos(parameters.Limit, parameters.Offset, parameters.OrderBy, parameters.Asc)

		if err != nil {
			return context.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return context.Status(fiber.StatusOK).JSON(todos)
	}
}

// Get Todo by Id
func Get(service Service) fiber.Handler {
	return func(context *fiber.Ctx) error {
		id := context.Params("id")

		if id == "" {
			return context.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "ID is required",
			})
		}

		todos, err := service.GetTodo(id)

		if err != nil {
			return context.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return context.Status(fiber.StatusOK).JSON(todos)
	}
}

// CreateTodo POST a new TODO item
func CreateTodo(service Service) fiber.Handler {
	return func(context *fiber.Ctx) error {
		// New TODO struct
		todoPayload := new(Todo)

		// Parse body into struct
		if err := context.BodyParser(todoPayload); err != nil {
			return context.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		err := todoPayload.Validate()
		if err != nil {
			// validationErrors := err.(validator.ValidationErrors)
			// fmt.Println(validationErrors)
			return context.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		todoInserted, err := service.InsertTodo(todoPayload)

		if err != nil {
			return context.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return context.Status(fiber.StatusOK).JSON(todoInserted)
	}
}
