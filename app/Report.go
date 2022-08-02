package app

import (
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
	reportFolderPath, err := filepath.Abs(config.Reportfolderpath)
	if err != nil {
		fmt.Println(err)
		return
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
		reportFile.WriteString("Time: " + uniqueException.Time)
		reportFile.WriteString(uniqueException.FirstLine + "\n\n")
		reportFile.WriteString(uniqueException.First10Lines + "\n")
		reportFile.WriteString("\n\n")
	}
}
