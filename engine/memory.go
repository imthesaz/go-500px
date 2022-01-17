package engine

var photoIDMemory map[string]bool

func InitPhotoIDMemory() {
	photoIDMemory = make(map[string]bool)
}

func AddIDToPhotoIDMap(id string) {
	photoIDMemory[id] = true
}

func CheckifIDExists(id string) bool {
	_, ok := photoIDMemory[id]
	return ok
}
