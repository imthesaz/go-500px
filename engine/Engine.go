package engine

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-500px/httpClient"
	"github.com/go-500px/models"
	"github.com/go-500px/utils"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
)

func Start(wg *sync.WaitGroup, config models.Config) {
	InitPhotoIDMemory()
	photoFileURLs := make(chan string, 20)

	go photoSearchGraphQL(wg, photoFileURLs, config)
	go photoDownload(wg, photoFileURLs, config)
}

func photoSearchGraphQL(wg *sync.WaitGroup, photoFileURLs chan<- string, config models.Config) {
	defer wg.Done()
	imageCounter := 0
	if config.SearchConfig.PhotoIndex != 0 {
		imageCounter = config.SearchConfig.PhotoIndex + 1
	}

	cursor, err := constructInitialCursorString(config.SearchConfig.PhotoIndex)
	if err != nil {
		log.Fatalln(err)
	}
	sortIDX := findSortStrIndex(config.SearchConfig.Sort, config.SearchConfig.PrevSort)
	photoSearchPaginationContainerQuery := &models.PhotoSearchPaginationContainerQuery{}

	for i := sortIDX; i < len(config.SearchConfig.Sort); i++ {
		hasNextPage := true
		downloadedCount := 0
		for {
			if hasNextPage && downloadedCount < config.SearchConfig.Count {

				photoSearchPaginationContainerQuery.InitPhotoSearchPaginationContainerQueryBody(cursor, config.SearchConfig.SearchTerm, config.SearchConfig.Sort[i])
				graphRes, err := httpClient.GetPhotoSearchPaginationContainer(photoSearchPaginationContainerQuery)
				if err != nil {
					log.Fatalln(err)
				}
				err = photoSearchDetail(graphRes, photoFileURLs, &imageCounter)
				if err != nil {
					log.Fatalln(err)
				}
				hasNextPage = graphRes.GetHasNextPage()
				cursor = graphRes.Data.PhotoSearch.PageInfo.EndCursor
				downloadedCount = imageCounter
			} else {
				break
			}

		}
	}
}

func photoDownload(wg *sync.WaitGroup, photoFileURLs <-chan string, config models.Config) {
	defer wg.Done()
	imageCounter := 0
	if config.SearchConfig.PhotoIndex != 0 {
		imageCounter = config.SearchConfig.PhotoIndex + 1
	}
	limit := 0
	if config.SearchConfig.Count < 10000 {
		limit = config.SearchConfig.Count
	} else {
		limit = 30000
	}

	for {
		if imageCounter != limit {
			url := <-photoFileURLs

			err := httpClient.DownloadPhotoFile(utils.FolderPath+"/"+strconv.Itoa(imageCounter)+".jpg", url)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(strconv.Itoa(imageCounter) + ".jpg downloaded")
			imageCounter++
		} else {
			break
		}

	}
}
func photoSearchDetail(graphQLRes *models.GraphQLResponse, photoFileURLs chan<- string, counter *int) error {
	nodeLength := len(graphQLRes.Data.PhotoSearch.Edges)
	batchCount := 0
	for i := 0; i < nodeLength; i++ {
		url := graphQLRes.Data.PhotoSearch.Edges[i].Node.Images[0].URL
		if !CheckifIDExists(url) {
			AddIDToPhotoIDMap(url)

			photoFileURLs <- url
			batchCount++
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

func findSortStrIndex(sort []string, sortStr string) int {
	for idx, str := range sort {
		if strings.Compare(str, sortStr) == 0 {
			return idx
		}

	}
	return 0
}

func constructInitialCursorString(cursor int) (string, error) {

	nextCursor := math.Ceil(float64(cursor)/20.0)*20 - 1
	if nextCursor >= 9999 {

		return "", errors.New(" the initial cursor value leads to a cursor for images higher than 10000")
	}
	return base64.StdEncoding.EncodeToString([]byte("pos-" + strconv.Itoa(int(nextCursor)))), nil

}
