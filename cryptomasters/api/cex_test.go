package api_test

import (
	"test/go/ccryptomasters/api"
	"testing"
)

func TestAPICall(t *testing.T)  {
	_,err := api.GetRate("")
	if err == nil {
		t.Errorf("Error not found")
	}
}