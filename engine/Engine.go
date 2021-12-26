package engine

import (
	"fmt"
	"github.com/go-500px/httpClient"
	"github.com/go-500px/models"
	"github.com/go-500px/utils"
	"log"
	"strconv"
	"sync"
)

func EngineStart(wg *sync.WaitGroup, searchStr string, sortStr string, limit int) {
	photoFileURLs := make(chan string)
	go photoSearchGraphQL(wg, searchStr, sortStr, photoFileURLs, limit)
	go photoDownload(wg, photoFileURLs, limit)
}

func photoSearchGraphQL(wg *sync.WaitGroup, searchStr string, sortStr string, photoFileURLs chan<- string, limit int) {
	defer wg.Done()
	imageCounter := 0
	photoSearchPaginationContainerQuery := &models.PhotoSearchPaginationContainerQuery{}
	photoSearchQueryRendererQuery := &models.PhotoSearchQueryRendererQuery{}
	photoSearchQueryRendererQuery.InitPhotoSearchQueryRendererQueryBody(searchStr, sortStr)

	graphRes, err := httpClient.GetPhotoSearchQueryRenderer(photoSearchQueryRendererQuery)
	if err != nil {
		log.Fatalln(err)
	}

	cursor := graphRes.Data.PhotoSearch.PageInfo.EndCursor
	err = photoSearchDetail(graphRes, photoFileURLs, &imageCounter)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		if graphRes.GetHasNextPage() && imageCounter < limit {

			photoSearchPaginationContainerQuery.InitPhotoSearchPaginationContainerQueryBody(cursor, searchStr, sortStr)
			graphRes, err := httpClient.GetPhotoSearchPaginationContainer(photoSearchPaginationContainerQuery)
			if err != nil {
				log.Fatalln(err)
			}
			err = photoSearchDetail(graphRes, photoFileURLs, &imageCounter)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			break
		}
	}
}

func photoDownload(wg *sync.WaitGroup, photoFileURLs <-chan string, limit int) {
	defer wg.Done()
	imageCounter := 0
	for {
		if imageCounter != limit {
			url := <-photoFileURLs

			err := httpClient.DownloadPhotoFile(utils.FolderPath+"/"+strconv.Itoa(imageCounter)+".jpg", url)
			if err != nil {
				log.Fatalln(err)
			}
			imageCounter++
		} else {
			break
		}

	}
}
func photoSearchDetail(graphQLRes *models.GraphQLResponse, photoFileURLs chan<- string, counter *int) error {
	nodeLength := len(graphQLRes.Data.PhotoSearch.Edges)
	for i := 0; i < nodeLength; i++ {
		photoFileURLs <- graphQLRes.Data.PhotoSearch.Edges[i].Node.Images[0].URL
		var detail []string
		extractGraphQLDetail(&detail, graphQLRes, i, counter)
		*counter++
		photoDetail, err := httpClient.GetPhotoDetail(graphQLRes.Data.PhotoSearch.Edges[i].Node.LegacyID)
		if err != nil {
			return err
		}
		extractPhotoDetail(&detail, photoDetail)
		err = utils.WritePhotoDetailRecord(detail)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractGraphQLDetail(detail *[]string, graphQLRes *models.GraphQLResponse, idx int, counter *int) {

	*detail = append(*detail, strconv.Itoa(*counter))
	*detail = append(*detail, graphQLRes.Data.PhotoSearch.Edges[idx].Node.Name)
	*detail = append(*detail, graphQLRes.Data.PhotoSearch.Edges[idx].Node.Description)
	*detail = append(*detail, graphQLRes.Data.PhotoSearch.Edges[idx].Node.Category)
	tags := ""
	for _, tag := range graphQLRes.Data.PhotoSearch.Edges[idx].Node.Tags {
		tags = tags + "$" + tag
	}
	*detail = append(*detail, tags)
}

func extractPhotoDetail(detail *[]string, photoDetail *models.PhotoDetail) {
	*detail = append(*detail, fmt.Sprintf("%f", photoDetail.Photos.PhotoInfo.Rating))
	*detail = append(*detail, photoDetail.Photos.PhotoInfo.TakenAt.String())
	*detail = append(*detail, photoDetail.Photos.PhotoInfo.ShutterSpeed)
	*detail = append(*detail, photoDetail.Photos.PhotoInfo.FocalLength)
	*detail = append(*detail, photoDetail.Photos.PhotoInfo.Aperture)
	*detail = append(*detail, photoDetail.Photos.PhotoInfo.Camera)
	*detail = append(*detail, photoDetail.Photos.PhotoInfo.Lens)
	*detail = append(*detail, photoDetail.Photos.PhotoInfo.Iso)
	*detail = append(*detail, photoDetail.Photos.PhotoInfo.Location)
	*detail = append(*detail, fmt.Sprintf("%f", photoDetail.Photos.PhotoInfo.Latitude))
	*detail = append(*detail, fmt.Sprintf("%f", photoDetail.Photos.PhotoInfo.Longitude))
	*detail = append(*detail, strconv.FormatBool(photoDetail.Photos.PhotoInfo.Nsfw))
}
