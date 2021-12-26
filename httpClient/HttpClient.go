package httpClient

import (
	"bytes"
	"encoding/json"
	"github.com/go-500px/models"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var mainQueryClient *http.Client

var photoDetailsClient *http.Client

func InitHTTPClients() {
	mainQueryClient = &http.Client{}
	photoDetailsClient = &http.Client{}
}

func GetPhotoSearchPaginationContainer(P *models.PhotoSearchPaginationContainerQuery) (*models.GraphQLResponse, error) {

	requestBody, err := json.Marshal(P)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", models.BaseGraphQLURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	resp, err := mainQueryClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	graphqlRes := &models.GraphQLResponse{}
	err = json.Unmarshal(body, graphqlRes)

	if err != nil {
		return nil, err
	}
	return graphqlRes, nil
}

func GetPhotoSearchQueryRenderer(P *models.PhotoSearchQueryRendererQuery) (*models.GraphQLResponse, error) {
	requestBody, err := json.Marshal(P)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", models.BaseGraphQLURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	resp, err := mainQueryClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	graphqlRes := &models.GraphQLResponse{}
	err = json.Unmarshal(body, graphqlRes)

	if err != nil {
		return nil, err
	}
	return graphqlRes, nil
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

func GetPhotoDetail(id string) (*models.PhotoDetail, error) {

	resp, err := photoDetailsClient.Get(models.CreatePhotoDetailQuery(id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonStr := strings.Replace(string(body), id, "photo_info", 1)

	photoDetail := &models.PhotoDetail{}
	err = json.Unmarshal([]byte(jsonStr), photoDetail)

	if err != nil {
		return nil, err
	}
	return photoDetail, nil
}
