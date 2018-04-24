package main 

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetBirdsHandler(t *testing.T)  {
	mockStore := InitMockStore()

	mockStore.On("GetBirds").Return([]*Bird{{"sparrow", "A small harmless bird"}},nil).Once()

	req,err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(GetBirdHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := Bird{"sparrow", "A small harmless bird"}
	b := []Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v and %v", actual)
	}
}