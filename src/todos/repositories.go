package todos

import (
	"database/sql"
	"fiber-estudos/src/databases"
	"fiber-estudos/src/utils"
	"fmt"
	"log"
)

// All users
func All(limit, offset int64) (todos []Todo, err error) {

	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	// create the select sql query
	query := `SELECT id, name, completed FROM todos LIMIT $1 OFFSET $2`

	rows, err := db.Query(query, limit, offset)
	defer rows.Close()
	utils.CheckError(err)

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Completed)
		utils.CheckError(err)

		todos = append(todos, todo)
	}

	return todos, nil
}

// Create todo
func Create(todo Todo) (Todo, error) {
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
func Get(id string) (Todo, error) {
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
