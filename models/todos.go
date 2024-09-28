package models

import (
	"github.com/gookit/color"
	"time"
)

type Todo struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	IsComplete bool      `json:"is_complete"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Todos []Todo

func (t *Todo) ColoredTitle() string {
	if t.IsComplete {
		return color.Green.Sprintf(t.Title)
	}

	var d int = int(time.Now().Sub(t.CreatedAt).Hours()) / 24
	switch {
	case d <= 1:
		return color.Green.Sprintf(t.Title)
	case d <= 3:
		return color.Yellow.Sprintf(t.Title)
	default:
		return color.Red.Sprintf(t.Title)
	}
}

func (t *Todo) ColoredStatus() string {
	switch t.IsComplete {
	case true:
		return color.Green.Sprintf("DONE")
	default:
		return color.Red.Sprintf("NOT DONE")
	}
}
