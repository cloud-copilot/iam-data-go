package iamdata

import (
	"testing"
)

func TestActionsForService(t *testing.T) {
	got, _ := ActionsForService("s3")

	if len(got) < 1 {
		t.Errorf("Returned action keys are empty")
	}
}

func TestActionsForServiceNotPresent(t *testing.T) {
	_, err := ActionsForService("r2")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestActionExistsTrue(t *testing.T) {
	got, _ := ActionExists("s3", "GetObject")
	want := true
	if got != want {
		t.Errorf("Expected s3 to have GetObject action")
	}
}

func TestActionExistsFalse(t *testing.T) {
	got, _ := ActionExists("s3", "GetSilliness")
	want := false
	if got != want {
		t.Errorf("Expected s3 to not have GetSilliness action")
	}
}

func TestActionExistsServiceNotPresent(t *testing.T) {
	_, err := ActionExists("r2", "GetObject")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestActionDetailsExists(t *testing.T) {
	details, _ := ActionDetails("s3", "getobject")
	got := details.Name
	want := "GetObject"
	if got != want {
		t.Errorf("Expected action name to be GetObject")
	}
}

func TestActionDetailsNotExists(t *testing.T) {
	_, err := ActionDetails("s3", "getsilliness")
	if err == nil {
		t.Errorf("Expected error looking for non-existent action")
	}
}

func TestActionDetailsServiceNotPresent(t *testing.T) {
	_, err := ActionDetails("r2", "GetObject")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
