package utils

import (
	"github.com/go-500px/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ParseSearchConfigFile(pathAddr string) (models.Config, error) {

	yfile, err := ioutil.ReadFile(pathAddr)

	if err != nil {
		return models.Config{}, err
	}

	data := models.Config{}

	err = yaml.Unmarshal(yfile, &data)

	if err != nil {
		return models.Config{}, err
	}

	return data, nil
}
