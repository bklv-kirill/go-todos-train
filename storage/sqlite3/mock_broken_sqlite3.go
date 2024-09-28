package sqlite3

import (
	"errors"
	"github.com/bklv-kirill/go-todos-train/models"
)

type MockBrokenStorage struct{}

func (ms *MockBrokenStorage) Get() (models.Todos, error) {
	return nil, errors.New("broken storage")
}

func (ms *MockBrokenStorage) Create(title string) error {
	return errors.New("broken storage")
}

func (ms *MockBrokenStorage) Delete(id int) error {
	return errors.New("broken storage")
}

func (ms *MockBrokenStorage) ChangeStatus(id int) error {
	return errors.New("broken storage")
}

func (ms *MockBrokenStorage) ChangeTitle(id int, title string) error {
	return errors.New("broken storage")
}
