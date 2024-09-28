package main

import (
	todos "github.com/bklv-kirill/go-todos-train"
	"github.com/bklv-kirill/go-todos-train/storage/sqlite3"
	"github.com/gookit/color"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	storage, err := sqlite3.NewStorage()
	if err != nil {
		color.Error.Tips(color.Yellow.Sprint(err.Error()))
		return
	}
	defer func() {
		_ = storage.Close()
	}()

	var commands *todos.Commands = todos.NewCommands(storage)
	if err = commands.Execute(); err != nil {
		color.Error.Tips(color.Yellow.Sprint(err.Error()))
	}
}
