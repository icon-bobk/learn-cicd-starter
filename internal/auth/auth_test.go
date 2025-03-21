package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyNoApiKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "erg")
	response, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("should error when api key does not exist %v", response)
	}

}
func TestGetAPIKeyGood(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey 1234")
	response, err := GetAPIKey(header)
	if err != nil {
		t.Errorf("should return no error %v", response)
	}

}
