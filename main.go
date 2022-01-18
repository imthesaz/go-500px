package main

import (
	"github.com/go-500px/engine"
	"github.com/go-500px/httpClient"
	"github.com/go-500px/utils"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	config, err := utils.ParseSearchConfigFile("config.yml")
	httpClient.InitHTTPClients()
	utils.CreateFilePath(config.SearchConfig.SearchTerm)
	err = utils.InitCSVWriter()
	if err != nil {
		log.Fatalln(err)
	}
	err = utils.CreateHistoryFile()
	if err != nil {
		log.Fatalln(err)
	}
	wg.Add(2)
	engine.Start(&wg, config)
	wg.Wait()
}
