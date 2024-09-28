package sqlite3

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/bklv-kirill/go-todos-train/models"
	"time"
)

var todoNotFound error = errors.New("todo not found")

type Storage struct {
	db *sql.DB
}

func NewStorage() (*Storage, error) {
	db, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		return nil, err
	}

	var storage *Storage = &Storage{
		db: db,
	}

	if err = storage.createTableIfNotExists(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *Storage) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) createTableIfNotExists() error {
	var q string = `
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(255) NOT NULL,
			is_completed BOOLEAN NOT NULL DEFAULT false,
			created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
		);`

	if _, err := s.db.Exec(q); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Get() (models.Todos, error) {
	var q string = "SELECT * FROM todos"

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	var todos models.Todos = models.Todos{}
	for rows.Next() {
		var todo models.Todo = models.Todo{}
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.IsComplete, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *Storage) Create(title string) error {
	var q string = "INSERT INTO todos (title) VALUES ($1)"

	if _, err := s.db.Exec(q, title); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(id int) error {
	if !s.Exist(id) {
		return todoNotFound
	}

	var q string = "DELETE FROM todos WHERE id = $1"

	if _, err := s.db.Exec(q, id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ChangeStatus(id int) error {
	todo, err := s.Find(id)
	if err != nil {
		return err
	}

	var q string = "UPDATE todos SET is_completed = $1, updated_at = $2  WHERE id = $3"
	if _, err := s.db.Exec(q, !todo.IsComplete, time.Now(), id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ChangeTitle(id int, title string) error {
	if !s.Exist(id) {
		return todoNotFound
	}

	var q string = "UPDATE todos SET title = $1, updated_at = $2 WHERE id = $3"
	if _, err := s.db.Exec(q, title, time.Now(), id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Find(id int) (models.Todo, error) {
	var todo models.Todo

	var row *sql.Row = s.db.QueryRow("SELECT * FROM todos WHERE id = $1 LIMIT 1", id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.IsComplete, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return todo, fmt.Errorf("todo with ID %d not found", id)
	}

	return todo, nil
}

func (s *Storage) Exist(id int) bool {
	var row *sql.Row = s.db.QueryRow("SELECT id FROM todos WHERE id = $1 LIMIT 1", id)
	if err := row.Scan(new(int)); err != nil {
		return false
	}

	return true
}
