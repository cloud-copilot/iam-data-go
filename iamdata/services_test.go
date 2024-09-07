package iamdata

import (
	"fmt"
	"testing"
)

func TestServiceKeys(t *testing.T) {
	got, _ := ServiceKeys()

	if len(got) < 1 {
		t.Errorf("Returned service keys are empty")
	}
}

func TestServiceNamePresent(t *testing.T) {
	got, _ := ServiceName("s3")
	want := "Amazon S3"
	if got != want {
		t.Errorf(fmt.Sprintf("Got: %s, Want: %s", got, want))
	}
}

func TestServiceNameNotPresent(t *testing.T) {
	_, err := ServiceName("r2")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestServiceExistsTrue(t *testing.T) {
	got, _ := ServiceExists("s3")
	want := true
	if got != want {
		t.Errorf(fmt.Sprintf("Got: %t, Want: %t", got, want))
	}
}

func TestServiceExistsFalse(t *testing.T) {
	got, _ := ServiceExists("r2")
	want := false
	if got != want {
		t.Errorf(fmt.Sprintf("Got: %t, Want: %t", got, want))
	}
}
