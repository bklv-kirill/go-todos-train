package go_todos_train

import (
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/bklv-kirill/go-todos-train/models"
	"github.com/gookit/color"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	list         string = "--ls"
	add                 = "--add"
	remove              = "--rm"
	changeStatus        = "--cs"
	changeTitle         = "--ct"
	help                = "--help"
)

const (
	somethingsWrong  string = "use --help for get additional info"
	invalidInputData        = "invalid input data. " + somethingsWrong
)

type Commands struct {
	s Storage
}

type Storage interface {
	Get() (models.Todos, error)
	Create(title string) error
	Delete(id int) error
	ChangeStatus(id int) error
	ChangeTitle(id int, title string) error
}

func NewCommands(s Storage) *Commands {
	return &Commands{
		s: s,
	}
}

func (c *Commands) Execute() (err error) {
	defer func() {
		if err != nil {
			return
		}
	}()

	if len(os.Args) < 2 {
		return errors.New(somethingsWrong)
	}

	var cmd string = os.Args[1]
	switch cmd {
	case list:
		err = c.list()
	case add:
		err = c.add()
	case remove:
		err = c.remove()
	case changeStatus:
		err = c.changeStatus()
	case changeTitle:
		err = c.changeTitle()
	case help:
		c.help()
	default:
		err = errors.New(somethingsWrong)
	}

	return err
}

func (c *Commands) list() error {
	todos, err := c.s.Get()
	if err != nil {
		return err
	}

	var t *simpletable.Table = simpletable.New()
	t.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{
				Align: simpletable.AlignCenter,
				Text:  "ID",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "TITLE",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "STATUS",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "CREATED AT",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "UPDATED AT",
			},
		},
	}

	for _, todo := range todos {
		r := []*simpletable.Cell{
			{
				Align: simpletable.AlignRight,
				Text:  fmt.Sprintf("%d", todo.ID),
			},
			{
				Align: simpletable.AlignCenter,
				Text:  fmt.Sprintf("%s", todo.ColoredTitle()),
			},
			{
				Align: simpletable.AlignCenter,
				Text:  fmt.Sprintf("%s", todo.ColoredStatus()),
			},
			{
				Align: simpletable.AlignCenter,
				Text:  fmt.Sprintf("%s", todo.CreatedAt.Format(time.RFC822)),
			},
			{
				Align: simpletable.AlignCenter,
				Text:  fmt.Sprintf("%s", todo.UpdatedAt.Format(time.RFC822)),
			},
		}

		t.Body.Cells = append(t.Body.Cells, r)
	}

	fmt.Println(t.String())

	return nil
}

func (c *Commands) add() error {
	if len(os.Args) < 3 {
		return errors.New(invalidInputData)
	}

	title, err := getTitle(2)
	if err != nil {
		return fmt.Errorf("%s %s", err.Error(), "for add new todo")
	}

	if err = c.s.Create(title); err != nil {
		return err
	}

	_ = c.list()

	return nil
}

func (c *Commands) remove() error {
	if len(os.Args) < 3 {
		return errors.New(invalidInputData)
	}

	id, err := getId()
	if err != nil {
		return fmt.Errorf("%s %s", err.Error(), "for delete todo")
	}

	if err = c.s.Delete(id); err != nil {
		return err
	}

	_ = c.list()

	return nil
}

func (c *Commands) changeStatus() error {
	if len(os.Args) < 3 {
		return errors.New(invalidInputData)
	}

	id, err := getId()
	if err != nil {
		return fmt.Errorf("%s %s", err.Error(), "for change todo status")
	}

	if err = c.s.ChangeStatus(id); err != nil {
		return err
	}

	_ = c.list()

	return nil
}

func (c *Commands) changeTitle() error {
	if len(os.Args) < 4 {
		return errors.New(invalidInputData)
	}

	id, err := getId()
	if err != nil {
		return fmt.Errorf("%s %s", err.Error(), "for change todo title")
	}

	title, err := getTitle(3)
	if err != nil {
		return fmt.Errorf("%s %s", err.Error(), "for change todo title")
	}

	if err = c.s.ChangeTitle(id, title); err != nil {
		return err
	}

	_ = c.list()

	return nil
}

func (c *Commands) help() {
	color.Green.Println("HELP:")
	fmt.Println("├───> " + color.Yellow.Sprint("--ls") + " -> for " + color.Cyan.Sprint("get") + " all todos")
	fmt.Println("├───> " + color.Yellow.Sprint("--add {title}") + " -> for " + color.Cyan.Sprint("add") + " new todo")
	fmt.Println("├───> " + color.Yellow.Sprint("--rm {id}") + " -> for " + color.Cyan.Sprint("delete") + " todo by ID")
	fmt.Println("├───> " + color.Yellow.Sprint("--cs {id}") + " -> for " + color.Cyan.Sprint("change") + " todo status by ID")
	fmt.Println("├───> " + color.Yellow.Sprint("--ct {id} {title}") + " -> for " + color.Cyan.Sprint("change") + " todo title by ID")
}

func getTitle(start int) (string, error) {
	var out []string = os.Args[start:]
	var title string = strings.Join(out, " ")
	title = trimString(title)

	var titleLen int = utf8.RuneCountInString(title)
	if titleLen == 0 || titleLen > 255 {
		return "", errors.New("invalid title")
	}

	return title, nil
}

func getId() (int, error) {
	var out string = os.Args[2]
	out = trimString(out)

	id, err := strconv.Atoi(out)
	if err != nil {
		return 0, errors.New("invalid ID")
	}

	return id, nil
}

func trimString(str string) string {
	str = strings.TrimSpace(str)
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\t")

	return str
}
