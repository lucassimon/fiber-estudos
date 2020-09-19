package todos

// Todo This is a model for TODO
type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
