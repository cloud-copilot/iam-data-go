package iamdata

import (
	"errors"
	"fmt"
	"strings"
)

// ConditionKeyType maps to the condition key metadata available.
type ConditionKeyType struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// ConditionKeysMap is a map of condition key names to ConditionKeyType metadata.
type ConditionKeysMap map[string]ConditionKeyType

// ConditionKeysForService returns a list of condition key names for a given service.
func ConditionKeysForService(serviceKey string) ([]string, error) {
	data, err := readConditionKeyData[ConditionKeysMap](serviceKey)
	if err != nil {
		return nil, err
	}

	var keys []string
	for _, condition := range data {
		keys = append(keys, condition.Key)
	}

	return keys, nil
}

// ConditionKeyExists checks if a condition key exists for a given service.
func ConditionKeyExists(serviceKey string, conditionKey string) (bool, error) {
	data, err := readConditionKeyData[ConditionKeysMap](serviceKey)
	if err != nil {
		return false, err
	}

	_, ok := data[strings.ToLower(conditionKey)]
	return ok, nil
}

// ConditionKeyDetails returns the metadata for a given condition key for a given service.
func ConditionKeyDetails(serviceKey string, conditionKey string) (ConditionKeyType, error) {
	data, err := readConditionKeyData[ConditionKeysMap](serviceKey)
	if err != nil {
		return ConditionKeyType{}, err
	}

	condition, ok := data[strings.ToLower(conditionKey)]
	if !ok {
		return ConditionKeyType{}, errors.New(fmt.Sprintf("Condition Key %s does not exist for service %s", conditionKey, serviceKey))
	}

	return condition, nil
}
