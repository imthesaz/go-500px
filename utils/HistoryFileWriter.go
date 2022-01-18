package utils

import (
	"bufio"
	"log"
	"os"
)

func CreateHistoryFile() error {
	exists, err := exists(FolderPath + "/history.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
		file, err := os.Create(FolderPath + "/history.txt")
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func AddPhotoIDRecord(id string) error {
	f, err := os.OpenFile(FolderPath+"/history.txt",
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(id + "\n"); err != nil {
		return err
	}
	return nil
}

func LoadHistoryID() ([]string, error) {
	file, err := os.Open(FolderPath + "/history.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var historyIDs []string

	for scanner.Scan() {
		historyIDs = append(historyIDs, scanner.Text())
	}
	return historyIDs, nil
}
