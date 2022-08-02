package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/goddamnnoob/loganalyzer/conf"
	"github.com/goddamnnoob/loganalyzer/exception"
)

func CreateNewReport(config *conf.Config, uniqueExceptions []exception.Exception) {
	currentTime := time.Now()
	reportFileName := "report_" + currentTime.Format("01-02-2006 15:04:05") + ".txt"
	fmt.Println(config.Reportsfolderpath)
	reportFolderPath, err := filepath.Abs(config.Reportsfolderpath)
	fmt.Println(reportFolderPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := os.Stat(reportFolderPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(reportFolderPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	reportFilePath := filepath.Join(reportFolderPath, reportFileName)
	reportFile, err := os.Create(reportFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	reportFile.WriteString(currentTime.String() + "\n")
	for _, uniqueException := range uniqueExceptions {
		reportFile.WriteString("-----" + uniqueException.Name + "[" + fmt.Sprint(uniqueException.Count) + "]" + "-----\n")
		reportFile.WriteString("Time: " + uniqueException.Time + "\n\n")
		reportFile.WriteString(uniqueException.FirstLine + "\n\n")
		reportFile.WriteString(uniqueException.First10Lines + "\n")
		reportFile.WriteString("\n\n")
	}
}
