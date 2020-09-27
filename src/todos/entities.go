package todos

// Todo This is a model for TODO
type Todo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
