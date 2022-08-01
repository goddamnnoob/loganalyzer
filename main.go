package main

import (
	"github.com/goddamnnoob/loganalyzer/app"
	"github.com/goddamnnoob/loganalyzer/conf"
)

func main() {
	config, err := conf.GetConfiguration()
	if err != nil {
		app.UniqueExceptions(config.Logsfolderpath)
	}
}
