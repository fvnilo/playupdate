package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	play_http "github.com/nylo-andry/playupdate/http"
)

func TestUpdate_ApiSucceeds(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(play_http.NewUpdateProfile())
		w.Write(b)
	}))

	service := play_http.NewPlayUpdateService(ts.URL)
	res, err := service.Update("foo:bar:baz")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}

	if res == nil {
		t.Error("Response shouldn't have been nil")
	}
}

func TestUpdate_ApiReturnsError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		e := &play_http.ApiError{
			StatusCode: http.StatusUnauthorized,
		}
		b, _ := json.Marshal(e)
		w.Write(b)
	}))

	service := play_http.NewPlayUpdateService(ts.URL)
	_, err := service.Update("foo:bar:baz")
	if err == nil {
		t.Error("Expected to have an error")
	}
	var v, ok = err.(play_http.ApiError)
	if !ok {
		t.Error("Returned error was not an ApiError")
	}

	if v.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected to have status code %v but got %v", http.StatusUnauthorized, v.StatusCode)
	}
}
