package engine

import "github.com/go-500px/utils"

var photoIDMemory map[string]bool

func InitPhotoIDMemory() error {
	photoIDMemory = make(map[string]bool)
	err := loadHistoryID()
	if err != nil {
		return err
	}
	return nil
}

func AddIDToPhotoIDMap(id string) error {
	photoIDMemory[id] = true
	err := addIDToHistory(id)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfIDExists(id string) bool {
	_, ok := photoIDMemory[id]
	return ok
}
func loadHistoryID() error {
	historyIDs, err := utils.LoadHistoryID()
	if err != nil {
		return err
	}
	for _, id := range historyIDs {
		AddIDToPhotoIDMap(id)
	}
	return nil
}

func addIDToHistory(id string) error {
	err := utils.AddPhotoIDRecord(id)
	if err != nil {
		return err
	}
	return nil
}
