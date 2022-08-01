package main

import (
	"fmt"

	"github.com/goddamnnoob/loganalyzer/app"
	"github.com/goddamnnoob/loganalyzer/conf"
)

func main() {
	config, err := conf.GetConfiguration("./config.json")
	if err == nil {
		fmt.Println(config)
		app.UniqueExceptions(config.Logsfolderpath)
	}
}
