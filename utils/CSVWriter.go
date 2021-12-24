package utils

import (
	"os"
	"path/filepath"
)

var file *os.File
var err error

func InitCSVWriter() error {
	file, err = os.Create(filepath.Join("result", filepath.Base("d.csv")))
	if err != nil {
		return err
	}
	return nil
}
