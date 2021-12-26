package utils

import (
	"encoding/csv"
	"log"
	"os"
)

var FolderPath string

func CreateFilePath(folderName string) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	err = os.Mkdir(path+"/Result/"+folderName, 0755)
	if err != nil {
		log.Println(err)
	}
	FolderPath = path + "/Result/" + folderName
}

func InitCSVWriter(folderName string) error {

	file, err := os.Create(FolderPath + "/details.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	initColumns := []string{"id", "name", "description", "category", "tags", "rating", "taken_at", "shutter_speed", "focal_length", "aperture", "camera", "Lens", "iso", "location", "latitude", "longitude", "nsfw"}
	csvwriter := csv.NewWriter(file)
	err = csvwriter.Write(initColumns)
	if err != nil {
		return err
	}
	csvwriter.Flush()
	return nil
}

func WritePhotoDetailRecord(detail []string) error {
	file, err := os.OpenFile(FolderPath+"/details.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	csvwriter := csv.NewWriter(file)
	err = csvwriter.Write(detail)
	if err != nil {
		return err
	}
	csvwriter.Flush()
	return nil
}
