# AWS IAM Data for Go

## Description
Contains IAM data for AWS actions, resources, and conditions based on IAM policy documents. This is intended to be used in downstream projects to provide a reference for IAM policy documents.

## Data Updates
Data is scanned daily and a new version is published if there are changes. The version number is updated to reflect the date of the last update and the function `UpdatedAt()` returns the date of the last data update. This process is managed outside this repo.

## Usage

```go
package main

import (
  "fmt"
  "github.com/cloud-copilot/iam-data-go/iamdata"
)

func main() {
  serviceKeys, _ := iamdata.ServiceKeys()
  for _, serviceKey := range serviceKeys {
    serviceName, _ := iamdata.ServiceName(serviceKey)
    fmt.Println("Getting Actions for", serviceName)
    actions, _ := iamdata.ActionsForService(serviceKey)
    for _, action := range actions {
      actionDetails, _ := iamdata.ActionDetails(serviceKey, action)
      fmt.Println(actionDetails)
    }
  }
}
```

## API
### Services
* `ServiceKeys()` - Returns a slice of all service keys such as 's3', 'ec2', etc.
* `ServiceName(serviceKey string)` - Returns the service name for a given service key.
* `ServiceExists(serviceKey string)` - Returns true if the service key exists.

### Actions
* `ActionsForService(serviceKey string)` - Returns an array of all actions for a given service key.
* `ActionDetails(serviceKey string actionName string)` - Returns an object with the action details such as `description`, `resourceTypes`, and `conditionKeys`.
* `ActionExists(serviceKey string, actionName string)` - Returns true if the action exists.

### Resource Types
* `ResourceTypesForService(serviceKey string)` - Returns an array of all resource types for a given service key.
* `ResourceTypesDetails(serviceKey string, resourceTypeKey string)` - Returns an object with the resource type details such as `description`, `arnFormat`, and `conditionKeys`.
* `ResourceTypesExists(serviceKey string resourceTypeKey: string)` - Returns true if the resource type exists.

### Conditions Keys
* `ConditionKeysForService(serviceKey string)` - Returns an array of all condition keys for a given service key.
* `ConditionKeyDetails(serviceKey string, conditionKey string)` - Returns an object with the condition key details such as `description`, `conditionValueTypes`, and `conditionOperators`.
* `ConditionKeyExists(serviceKey string, conditionKey string)` - Returns true if the condition key exists.

### Version Info
The version is numper is formatted as `major.minor.updatedAt`. The updatedAt is the date the data was last updated in the format `YYYYMMDDX` where `X` is a counter to enable publishing more than once per day if necessary. For example version `0.1.202408291` has data updated on August 29th, 2024.

The version can be accessed using the `Version()` method.

There is also `UpdatedAt()` which returns the date the data was last updated.
