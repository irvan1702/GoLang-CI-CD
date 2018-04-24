package main

import(
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
)

func TestHandlers(t *testing.T)  {
	//create router
	r := newRouter()

	//create new server with httptest library 'NewServer' method
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d",
		resp.StatusCode)
	}

	//execute close until the func is return something
	defer resp.Body.Close()

	//read body by bytes
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	//convert bytes to string
	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", 
			expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T){
	r := newRouter()
	mockServer := httptest.NewServer(r)
	
	resp, err := http.Post(mockServer.URL+"/hello","",nil)

	if err != nil {
		t.Fatal(err)
	}

	//405 Method not allowed
	if resp.StatusCode != http.StatusMethodNotAllowed{
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T)  {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL+ "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK{
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType{
		t.Errorf("Wrong Content type, expected %s, got %s",expectedContentType,contentType)
	}
}
