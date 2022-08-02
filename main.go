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
	fmt.Println(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	uniqueExceptions = app.UniqueExceptions(config)
	fmt.Println(uniqueExceptions)
	if uniqueExceptions != nil {
		app.CreateNewReport(config, uniqueExceptions)
	}
}
