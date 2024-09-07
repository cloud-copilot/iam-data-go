package iamdata

import (
	"fmt"
	"time"
)

// Metadata available for the module.
type Metadata struct {
	Version   string    `json:"version"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// metadata returns the metadata for the module.
func metadata() (Metadata, error) {
	data, err := readJSONData[Metadata]([]string{"metadata"})
	if err != nil {
		return Metadata{}, err
	}
	return data, nil
}

// UpdatedAt returns the time the IAM data was last updated.
func UpdatedAt() (time.Time, error) {
	data, err := metadata()
	if err != nil {
		fmt.Println("Error getting metadata:", err)
		return time.Time{}, err
	}

	return data.UpdatedAt, nil
}

// Version returns the version of the IAM data module.
func Version() (string, error) {
	data, err := metadata()
	if err != nil {
		fmt.Println("Error getting metadata:", err)
		return "", err
	}
	return data.Version, nil
}
