package iamdata

import (
	"testing"
)

func TestConditionKeysForService(t *testing.T) {
	got, _ := ConditionKeysForService("s3")

	if len(got) < 1 {
		t.Errorf("Returned action keys are empty")
	}
}

func TestConditionKeysForServicePresent(t *testing.T) {
	_, err := ConditionKeysForService("r2")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestConditionKeyExistsTrue(t *testing.T) {
	got, _ := ConditionKeyExists("s3", "s3:prefix")
	want := true
	if got != want {
		t.Errorf("Expected s3 to have s3:prefix condition key")
	}
}

func TestConditionKeyExistsFalse(t *testing.T) {
	got, _ := ConditionKeyExists("s3", "s3:silliness")
	want := false
	if got != want {
		t.Errorf("Expected s3 to not have s3:silliness condition key")
	}
}

func TestConditionKeyExistsServiceNotPresent(t *testing.T) {
	_, err := ConditionKeyExists("r2", "s3:prefix")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestConditionKeyDetailsExists(t *testing.T) {
	details, _ := ConditionKeyDetails("s3", "aws:TagKeys")
	got := details.Key
	want := "aws:TagKeys"
	if got != want {
		t.Errorf("Expected condition key to be aws:TagKeys")
	}
}

func TestConditionKeyDetailsNotExists(t *testing.T) {
	_, err := ConditionKeyDetails("s3", "s3:silliness")
	if err == nil {
		t.Errorf("Expected error looking for non-existent condition key")
	}
}

func TestConditionKeyDetailsServiceNotPresent(t *testing.T) {
	_, err := ConditionKeyDetails("r2", "s3:prefix")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}