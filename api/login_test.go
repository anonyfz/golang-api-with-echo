package api_test

import (
	"reflect"
	"testing"

	. "go-echo-api/api"
)

func TestNewAPI(t *testing.T) {
	service := MockLoginAPI{}
	expectedLoginAPI := &LoginAPI{
		Service: service,
	}

	actualLoginAPI := NewLoginAPI(service)

	if !reflect.DeepEqual(actualLoginAPI, expectedLoginAPI) {
		t.Errorf("expect %v but it got %v", expectedLoginAPI, actualLoginAPI)
	}
}
