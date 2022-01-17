package engine

import (
	"errors"
	"fmt"
	"github.com/go-500px/httpClient"
	"github.com/go-500px/models"
	"github.com/go-500px/utils"
	"log"
	"strconv"
	"sync"
)

func EngineStart(wg *sync.WaitGroup, searchStr string, sortStr string, limit int) {
	InitPhotoIDMemory()
	photoFileURLs := make(chan string)
	downloadConfirm := make(chan int)
	go photoSearchGraphQL(wg, searchStr, sortStr, photoFileURLs, downloadConfirm, limit)
	go photoDownload(wg, photoFileURLs, downloadConfirm, limit)
}

func photoSearchGraphQL(wg *sync.WaitGroup, searchStr string, sortStr string, photoFileURLs chan<- string, downloadConfirm <-chan int, limit int) {
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
	err = photoSearchDetail(graphRes, photoFileURLs, downloadConfirm, &imageCounter)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		if graphRes.GetHasNextPage() && imageCounter < limit {

			photoSearchPaginationContainerQuery.InitPhotoSearchPaginationContainerQueryBody(cursor, searchStr, sortStr)
			graphRes, err = httpClient.GetPhotoSearchPaginationContainer(photoSearchPaginationContainerQuery)
			if err != nil {
				log.Fatalln(err)
			}
			err = photoSearchDetail(graphRes, photoFileURLs, downloadConfirm, &imageCounter)
			if err != nil {
				log.Fatalln(err)
			}
			cursor = graphRes.Data.PhotoSearch.PageInfo.EndCursor
		} else {
			break
		}

	}
}

func photoDownload(wg *sync.WaitGroup, photoFileURLs <-chan string, downloadConfirm chan<- int, limit int) {
	defer wg.Done()
	imageCounter := 0
	batchCounter := 0
	var confirm int
	confirm = 1
	for {
		if imageCounter != limit {
			url := <-photoFileURLs

			err := httpClient.DownloadPhotoFile(utils.FolderPath+"/"+strconv.Itoa(imageCounter)+".jpg", url)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(strconv.Itoa(imageCounter) + ".jpg downloaded")
			imageCounter++
			batchCounter++
			if batchCounter == 20 {
				batchCounter = 0
				downloadConfirm <- confirm
			}
		} else {
			break
		}

	}
}
func photoSearchDetail(graphQLRes *models.GraphQLResponse, photoFileURLs chan<- string, downloadConfirm <-chan int, counter *int) error {
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
	confirm := <-downloadConfirm
	if confirm == 1 {
		return nil
	} else {
		return errors.New("problem in batch download")
	}
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
