package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoList(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t TodoListPostgres) CreateTodoList(userId int, list todo.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createTodoListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListsTable)
	row := t.db.QueryRow(createTodoListQuery, list.Title, list.Description)

	if err := row.Scan(&id); err != nil {
		err := tx.Rollback()
		return 0, err
	}

	createTodoListUserQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2)", usersListsTable)
	_, err = tx.Exec(createTodoListUserQuery, userId, id)
	if err != nil {
		err := tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
