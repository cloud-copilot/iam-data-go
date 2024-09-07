package iamdata

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed data/*.json data/**/*.json
var dataFiles embed.FS

// readJSONData reads a JSON file from the data folder and unmarshals it into a variable.
func readJSONData[T any](filePath []string) (T, error) {

	// Add the data folder at the beginning of the file path
	filePath = append([]string{"data"}, filePath...)
	// Add the .json extension to the last element of the file path
	lastIndex := len(filePath) - 1
	filePath[lastIndex] = filePath[lastIndex] + ".json"

	//Read the JSON file
	fullPath := filepath.Join(filePath...)
	// jsonData, err := fs.ReadFile(dataFiles, "data/" + filePath)
	jsonData, err := fs.ReadFile(dataFiles, fullPath)
	if err != nil {
		return *new(T), fmt.Errorf("error reading file: %w", err)
	}

	var data T
	// Unmarshal JSON into the variable
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return *new(T), fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return data, nil
}

// readServiceData reads the service data from the JSON file for the given service key.
func readActionData[T any](serviceKey string) (T, error) {
	data, err := readJSONData[T]([]string{"actions", strings.ToLower(serviceKey)})
	if err != nil {
		return *new(T), err
	}
	return data, nil
}

// readConditionKeyData reads the condition key data from the JSON file for the given service key.
func readConditionKeyData[T any](serviceKey string) (T, error) {
	data, err := readJSONData[T]([]string{"conditionKeys", strings.ToLower(serviceKey)})
	if err != nil {
		return *new(T), err
	}
	return data, nil
}

// readResourceTypeData reads the resource type data from the JSON file for the given service key.
func readResourceTypeData[T any](serviceKey string) (T, error) {
	data, err := readJSONData[T]([]string{"resourceTypes/", strings.ToLower(serviceKey)})
	if err != nil {
		return *new(T), err
	}
	return data, nil
}
