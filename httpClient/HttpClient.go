package httpClient

import (
	"bytes"
	"encoding/json"
	"github.com/go-500px/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var mainQueryClient *http.Client

var photoDetailsClient *http.Client

func GetPhotoSearchPaginationContainer(P *models.PhotoSearchPaginationContainerQuery) {
	mainQueryClient := &http.Client{}
	requestBody, err := json.Marshal(P)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest("POST", models.BaseGraphQLURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-type", "application/json")
	resp, err := mainQueryClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func GetPhotoSearchQueryRenderer(P *models.PhotoSearchQueryRendererQuery) {

}

func DownloadPhotoFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
