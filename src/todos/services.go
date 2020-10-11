package todos

//Service is an interface from which our api module can access our repository of all our models
type Service interface {
	GetTodo(id string) (Todo, error)
	FetchTodos(limit, offset int64, orderby string, asc string) ([]Todo, error)
	InsertTodo(todo *Todo) (Todo, error)
	// UpdateTodo(todo *Todo) (*Todo, error)
	// RemoveTodo(ID string) error
}

type service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertTodo(todo *Todo) (Todo, error) {
	return s.repository.CreateTodo(todo)
}

func (s *service) GetTodo(id string) (Todo, error) {
	return s.repository.Get(id)
}

func (s *service) FetchTodos(limit, offset int64, orderby string, asc string) ([]Todo, error) {
	return s.repository.FetchAllTodo(limit, offset, orderby, asc)
}

// func (s *service) UpdateTodo(book *Todo) (*Todo, error) {
// 	return s.repository.UpdateBook(todo)
// }

// func (s *service) RemoveTodo(ID string) error {
// 	return s.repository.DeleteBook(ID)
// }
