package main

import (
	"github.com/goddamnnoob/loganalyzer/app"
)

func main() {
	config, err := GetConfiguration()
	if err != nil {
		app.UniqueExceptions(config.Logsfolderpath)
	}
}
