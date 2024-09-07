package iamdata

import (
	"errors"
	"fmt"
	"strings"
)

// ServiceNameMap maps service keys to service names.
type ServiceNameMap map[string]string

// ServiceNames is a list of service keys.
type ServiceKeyStrings []string

// ServiceKeys returns a list of all IAM service keys.
func ServiceKeys() ([]string, error) {
	data, err := readJSONData[ServiceKeyStrings]([]string{"services"})
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ServiceName returns the name of a service given its key.
func ServiceName(serviceKey string) (string, error) {
	var data ServiceNameMap
	data, err := readJSONData[ServiceNameMap]([]string{"serviceNames"})
	if err != nil {
		return "", err
	}

	serviceName, ok := data[strings.ToLower(serviceKey)]
	if !ok {
		return "", errors.New(fmt.Sprintf("Service %s does not exist", serviceKey))
	}

	return serviceName, nil
}

// ServiceExists checks if a service exists.
func ServiceExists(serviceName string) (bool, error) {
	var data ServiceNameMap
	data, err := readJSONData[ServiceNameMap]([]string{"serviceNames"})
	if err != nil {
		return false, err
	}

	_, ok := data[strings.ToLower(serviceName)]
	return ok, nil
}
