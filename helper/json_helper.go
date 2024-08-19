package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Encoding data to JSON file
func SaveToJSON(fileName string, key interface{}) error {
	mu := &sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v\n", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	encodeJSON := json.NewEncoder(file)
	err = encodeJSON.Encode(key)
	if err != nil {
		return fmt.Errorf("error writing to JSON file: %v", err)
	}

	return nil
}

// Decoding the data of a JSON file
func LoadFromJSON(fileName string, key interface{}) error {
	mu := &sync.RWMutex{}
	mu.RLock()
	defer mu.RUnlock()

	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v\n", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	// I think the bug that removes the entire json is here

	// Check if the file is empty
	fileInfo, eErr := file.Stat()
	if eErr != nil {
		return fmt.Errorf("Failed to stat file: %v\n", eErr)
	}

	if fileInfo.Size() == 0 {
		// If the file is empty, return nil (no error) as there is nothing to load
		return nil
	}

	decodeJSON := json.NewDecoder(file)
	wErr := decodeJSON.Decode(key)
	if wErr != nil {
		return fmt.Errorf("error reading from JSON file: %v", wErr)
	}

	return nil
}
