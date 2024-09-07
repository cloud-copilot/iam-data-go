package iamdata

import (
	"testing"
)

func TestResourceTypesForService(t *testing.T) {
	got, _ := ResourceTypesForService("s3")

	if len(got) < 1 {
		t.Errorf("Returned action keys are empty")
	}
}

func TestResourceTypesForServiceNotPresent(t *testing.T) {
	_, err := ResourceTypesForService("r2")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestResourceTypesExistsTrue(t *testing.T) {
	got, _ := ResourceTypesExists("s3", "bucket")
	want := true
	if got != want {
		t.Errorf("Expected s3 to have bucket resource type")
	}
}

func TestResourceTypesExistsFalse(t *testing.T) {
	got, _ := ResourceTypesExists("s3", "silliness")
	want := false
	if got != want {
		t.Errorf("Expected s3 to not have silliness resource type")
	}
}

func TestResourceTypesExistsServiceNotPresent(t *testing.T) {
	_, err := ResourceTypesExists("r2", "bucket")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestResourceTypesDetailsExists(t *testing.T) {
	details, _ := ResourceTypesDetails("s3", "bucket")
	got := details.Key
	want := "bucket"
	if got != want {
		t.Errorf("Expected resource type to be bucket")
	}
}

func TestResourceTypesDetailsNotExists(t *testing.T) {
	_, err := ResourceTypesDetails("s3", "silliness")
	if err == nil {
		t.Errorf("Expected error looking for non-existent resource type")
	}
}

func TestResourceTypesServiceNotPresent(t *testing.T) {
	_, err := ResourceTypesDetails("r2", "bucket")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}