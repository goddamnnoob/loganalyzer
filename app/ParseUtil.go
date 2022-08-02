package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/goddamnnoob/loganalyzer/exception"
)

var uniqueExceptions []exception.Exception

func parseServerOut(filePath *string) ([]exception.Exception, *error) {
	file, e := os.Open(*filePath)
	lineIndex := 1
	if e != nil {
		fmt.Println(e)
		return nil, &e
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Exception:") {
			currentexception := exceptionPostProcessor(scanner)
			lineIndex += 9
			if isUniqueException(currentexception) {
				uniqueExceptions = append(uniqueExceptions, *currentexception)
			}
		}
		lineIndex++
	}
	return uniqueExceptions, nil
}

func exceptionPostProcessor(scanner *bufio.Scanner) *exception.Exception {
	var exception exception.Exception
	firstLine := scanner.Text()
	time := firstLine[0:27]
	name := getExceptionName(firstLine)
	exception.FirstLine = firstLine
	exception.Time = time
	exception.Name = name
	var first10Lines string
	for i := 0; i < 20; i++ {
		first10Lines += scanner.Text() + "\n"
		scanner.Scan()
	}
	exception.First10Lines = first10Lines
	return &exception
}

func getExceptionName(firstline string) string {
	var name string
	indexStart := strings.Index(firstline, "Exception")
	index := indexStart
	for index > 0 && string(firstline[index]) != ":" {
		index--
	}
	for index < indexStart {
		name += string(firstline[index])
		index++
	}
	name += "Exception"
	return name
}

func isUniqueException(exception *exception.Exception) bool {
	for _, ex := range uniqueExceptions {
		if strings.Compare(exception.Name, ex.Name) == 0 {
			exception.Count++
			return false
		}
	}
	return true
}
