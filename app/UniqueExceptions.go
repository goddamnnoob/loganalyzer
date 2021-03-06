package app

import (
	"fmt"
	"os"
	"strings"
)

func UniqueExceptions(logsFolderPath *string) {
	//serverOutFilesCount := 0
	if !isValidPath(logsFolderPath) {
		fmt.Println("Invalid Logs Folder Path specified !!!!!!")
		return
	}
	if !isValidDirectory(logsFolderPath) {
		fmt.Println("Logs Folder Path is not a Directory !!!!!!")
		return
	}
	filesInDirectory := getFilesListInFolder(logsFolderPath)
	serverOutFilesInDirectory := getServerOutFiles(filesInDirectory)
	fmt.Println(len(*serverOutFilesInDirectory))
}

func isValidPath(path *string) bool {
	_, err := os.Open(*path)
	return err == nil
}

func isValidDirectory(path *string) bool {
	dir, _ := os.Stat(*path)
	return dir.IsDir()
}

func getFilesListInFolder(path *string) *[]string {
	var filesInDirectory []string
	file, _ := os.Open(*path)
	filesList, _ := file.Readdir(0)
	for _, f := range filesList {
		filesInDirectory = append(filesInDirectory, f.Name())
	}
	return &filesInDirectory
}

func getServerOutFiles(filesInDirectory *[]string) *[]string {
	serverOut := "serverOut_"
	var serverOutFiles []string
	for _, f := range *filesInDirectory {
		if strings.Contains(f, serverOut) {
			serverOutFiles = append(serverOutFiles, f)
		}
	}
	return &serverOutFiles
}
