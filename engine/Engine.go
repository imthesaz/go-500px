package engine

import (
	"github.com/go-500px/httpClient"
	"github.com/go-500px/models"
	"log"
)

func EngineStart(searchStr string, sortStr string) {
	go photoSearchGraphQL(searchStr, sortStr)
	go photoDownload()
}

func photoSearchGraphQL(searchStr string, sortStr string) {

	photoSearchQueryRendererQuery := &models.PhotoSearchQueryRendererQuery{}
	photoSearchQueryRendererQuery.InitPhotoSearchQueryRendererQueryBody(searchStr, sortStr)

	graphRes, err := httpClient.GetPhotoSearchQueryRenderer(photoSearchQueryRendererQuery)

	if err != nil {
		log.Panic(err)
	}

	for {
		if graphRes.GetHasNextPage() {

		}
	}
}

func photoDownload() {

}
func photoSearchDetail() {

}
