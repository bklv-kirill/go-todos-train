package sqlite3

import (
	"errors"
	"github.com/bklv-kirill/go-todos-train/models"
	"time"
)

type MockStorage struct{}

func (ms *MockStorage) Get() (models.Todos, error) {
	return models.Todos{
		models.Todo{
			ID:         1,
			Title:      "Test todo",
			IsComplete: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}, nil
}

func (ms *MockStorage) Create(title string) error {
	if title == "error" {
		return errors.New("broken storage")
	}

	return nil
}

func (ms *MockStorage) Delete(id int) error {
	if id == 999 {
		return errors.New("broken storage")
	}

	return nil
}

func (ms *MockStorage) ChangeStatus(id int) error {
	if id == 999 {
		return errors.New("broken storage")
	}

	return nil
}

func (ms *MockStorage) ChangeTitle(id int, title string) error {
	if id == 999 {
		return errors.New("broken storage")
	}

	return nil
}
