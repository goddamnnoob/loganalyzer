package main

import (
	"fmt"

	"github.com/goddamnnoob/loganalyzer/app"
	"github.com/goddamnnoob/loganalyzer/conf"
	"github.com/goddamnnoob/loganalyzer/exception"
)

func main() {
	var uniqueExceptions []exception.Exception
	config, err := conf.GetConfiguration()
	if err != nil {
		fmt.Println(err)
		return
	}
	uniqueExceptions = app.UniqueExceptions(config)
	if uniqueExceptions != nil {
		app.CreateNewReport(config, uniqueExceptions)
	}
}
