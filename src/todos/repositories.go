package todos

import (
	"database/sql"
	"fiber-estudos/src/databases"
	"fiber-estudos/src/utils"
	"fmt"
	"log"
)

//Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateTodo(todo *Todo) (Todo, error)
	FetchAllTodo(limit, offset int64, orderby string, asc string) ([]Todo, error)
	Get(id string) (Todo, error)
	// UpdateTodo(todo *Todo) (*Todo, error)
	// DeleteTodo(ID string) error
}

type repository struct{}

//NewRepo is the single instance repo that is being created.
func NewRepo() Repository {
	return &repository{}
}

// FetchAllTodo All todos
func (r *repository) FetchAllTodo(limit, offset int64, orderBy string, asc string) (todos []Todo, err error) {
	if limit == 0 {
		limit = 10
	}

	if orderBy == "" {
		orderBy = "created"
	}

	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	query := `SELECT id, name, completed, created_at FROM todos ORDER BY $1 ASC LIMIT $2 OFFSET $3`

	// create the select sql query
	if asc == "0" {
		query = `SELECT id, name, completed, created_at FROM todos ORDER BY $1 DESC LIMIT $2 OFFSET $3`
	}

	fmt.Println(query)

	rows, err := db.Query(query, orderBy, limit, offset)
	defer rows.Close()
	utils.CheckError(err)

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		utils.CheckError(err)

		todos = append(todos, todo)
	}

	return todos, nil
}

// CreateTodo
func (r *repository) CreateTodo(todo *Todo) (Todo, error) {
	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()
	// the inserted id will store in this id

	var todoInserted Todo
	// hardcoded
	stmt := `INSERT INTO todos (name, completed) VALUES ($1, $2) RETURNING id, name, completed`

	// execute the sql statement
	// Scan function will save the insert id in the id
	rows, err := db.Query(stmt, todo.Name, todo.Completed)
	defer rows.Close()
	if err != nil {
		return todoInserted, err
	}

	log.Println("Inserted a single record", rows)

	for rows.Next() {
		err = rows.Scan(&todoInserted.ID, &todoInserted.Name, &todoInserted.Completed)
		utils.CheckError(err)
	}

	// return the inserted id
	return todoInserted, nil
}

// Get Fetch todo by id
func (r *repository) Get(id string) (Todo, error) {
	// hardcoded
	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	var todo Todo

	// create the select sql query
	stmt := `SELECT id, name, completed FROM todos WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(stmt, id)

	// unmarshal the row object to user
	err := row.Scan(&todo.ID, &todo.Name, &todo.Completed)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return todo, nil
	case nil:
		return todo, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return todo, err
}
