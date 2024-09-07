package iamdata

import (
	"strings"
	"errors"
	"fmt"
)

// ActionResourceType maps to the resource type metadata that can be included in an action.
type ActionResourceType struct {
	Name             string    `json:"name"`
	Required         bool      `json:"required"`
	ConditionKeys    []string  `json:"conditionKeys"`
	DependentActions []string  `json:"dependentActions"`
}

// AccessLevel is a custom type to restrict the values to specific strings
type AccessLevel string
const (
	Read    AccessLevel = "Read"
	Write   AccessLevel = "Write"
	List    AccessLevel = "List"
	Tagging AccessLevel = "Tagging"
)

// Action maps to the Action metadata available.
type Action struct {
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	AccessLevel      AccessLevel          `json:"accessLevel"`
	ResourceTypes    []ActionResourceType `json:"resourceTypes"`
	ConditionKeys    []string             `json:"conditionKeys"`
	DependentActions []string             `json:"dependentActions"`
}

// Actions is a map of action names to Action metadata.
type Actions map[string]Action

// ActionsForService returns a list of action names for a given service.
func ActionsForService(serviceKey string) ([]string, error) {
	data, err := readActionData[Actions](serviceKey)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, action := range data {
			names = append(names, action.Name)
	}
	return names, nil
}

// ActionExists checks if an action exists for a given service.
func ActionExists(serviceKey string, actionName string) (bool, error) {
	data, err := readActionData[Actions](serviceKey)
	if err != nil {
		return false, err
	}

	_, ok := data[strings.ToLower(actionName)]
	return ok, nil
}

// ActionDetails returns the metadata for a given action within a service.
func ActionDetails(serviceKey string, actionName string) (Action, error) {
	data, err := readActionData[Actions](serviceKey)
	if err != nil {
		return Action{}, err
	}

	action, ok := data[strings.ToLower(actionName)]
	if !ok {
		return Action{}, errors.New(fmt.Sprintf("Action %s does not exist for service %s", actionName, serviceKey))
	}
	return action, nil
}
