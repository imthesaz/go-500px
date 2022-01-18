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
	exists, err := exists(path + "/Result/" + folderName)
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		err = os.Mkdir(path+"/Result/"+folderName, 0755)
		if err != nil {
			log.Println(err)
		}
	}
	FolderPath = path + "/Result/" + folderName
}

func InitCSVWriter() error {
	exists, err := exists(FolderPath + "/details.csv")
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
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
	}
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

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
