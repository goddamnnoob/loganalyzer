package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goddamnnoob/loganalyzer/conf"
	"github.com/goddamnnoob/loganalyzer/exception"
)

func UniqueExceptions(config *conf.Config) []exception.Exception {
	var uniqueExceptions []exception.Exception
	logsFolderPath, e := filepath.Abs(config.Logsfolderpath)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	if !isValidPath(logsFolderPath) {
		fmt.Println("Invalid Logs Folder Path specified !!!!!! LogsFolderPath:" + logsFolderPath)
		return nil
	}
	if !isValidDirectory(logsFolderPath) {
		fmt.Println("Logs Folder Path is not a Directory !!!!!!")
		return nil
	}
	filesInDirectory := getFilesListInFolder(logsFolderPath)
	serverOutFilesInDirectory := getServerOutFiles(filesInDirectory)
	for _, serverOutFilePath := range serverOutFilesInDirectory {
		fmt.Println(serverOutFilePath)
		batchUniqueExceptions, err := parseServerOut(&serverOutFilePath)
		if err != nil {
			uniqueExceptions = append(uniqueExceptions, batchUniqueExceptions...)
		}
	}
	return uniqueExceptions
}

func isValidPath(path string) bool {
	file, err := os.Open(path)
	defer file.Close()
	return err == nil
}

func isValidDirectory(path string) bool {
	dir, _ := os.Stat(path)
	return dir.IsDir()
}

func getFilesListInFolder(path string) []string {
	var filesInDirectory []string
	file, _ := os.Open(path)
	filesList, _ := file.Readdir(0)
	for _, f := range filesList {
		filesInDirectory = append(filesInDirectory, filepath.Join(path, f.Name()))
	}
	defer file.Close()
	return filesInDirectory
}

func getServerOutFiles(filesInDirectory []string) []string {
	serverOut := "serverOut_"
	var serverOutFiles []string
	for _, f := range filesInDirectory {
		if strings.Contains(f, serverOut) {
			serverOutFiles = append(serverOutFiles, f)
		}
	}
	return serverOutFiles
}
