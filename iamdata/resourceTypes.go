package iamdata

import (
	"errors"
	"fmt"
	"strings"
)

// ResourceType maps to the resource type metadata available.
type ResourceType struct {
	Key           string    `json:"key"`
	Arn           string    `json:"arn"`
	ConditionKeys []string  `json:"conditionKeys"`
}

// ResourceTypesMap is a map of resource type names to ResourceType metadata.
type ResourceTypesMap map[string]ResourceType

// ResourceTypesForService returns a list of resource type names for a given service.
func ResourceTypesForService(serviceKey string) ([]string, error) {
	data, err := readResourceTypeData[ResourceTypesMap](serviceKey)
	if err != nil {
		return nil, err
	}

	var keys []string
	for _, resourceType := range data {
		keys = append(keys, resourceType.Key)
	}

	return keys, nil
}

// ResourceTypesExists checks if a resource type exists for a given service.
func ResourceTypesExists(serviceKey string, resourceTypeKey string) (bool, error) {
	data, err := readResourceTypeData[ResourceTypesMap](serviceKey)
	if err != nil {
		return false, err
	}

	_, ok := data[strings.ToLower(resourceTypeKey)]
	return ok, nil
}

// ResourceTypesDetails returns the metadata for a given resource type for a given service.
func ResourceTypesDetails(serviceKey string, resourceTypeKey string) (ResourceType, error) {
	data, err := readResourceTypeData[ResourceTypesMap](serviceKey)
	if err != nil {
		return ResourceType{}, err
	}

	resourceType, ok := data[strings.ToLower(resourceTypeKey)]
	if !ok {
		return ResourceType{}, errors.New(fmt.Sprintf("Resource Type %s does not exist for service %s", resourceTypeKey, serviceKey))
	}

	return resourceType, nil
}