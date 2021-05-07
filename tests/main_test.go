package tests

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	v1 "github.com/avkosme/golang-api-boilerplate/internal/delivery/http/v1"
	"github.com/julienschmidt/httprouter"
)

const readerJsonSuccess string = `{"update_id":899835778,"message":{"message_id":111,"from":{"id":111111111,"is_bot":false,"first_name":"Richi","last_name":"Blackman","username":"godzilla","language_code":"en"},"chat":{"id":22222222,"first_name":"Richi","last_name":"Blackman","username":"godzilla","type":"private"},"date":1618634228,"text":"/pages","entities":[{"offset":0,"length":6,"type":"bot_command"}]}}`

const readerJsonError string = `{"update_id":899835778,"message":{"message_id":111,"from":{"id":111111111,"is_bot":false,"first_name":"Richi","last_name":"Blackman","username":"godzilla","language_code":"en"},"chat":{"id":22222222,"first_name":"Richi","last_name":"Blackman","username":"godzilla","type":"private"},"date":1618634228,"text":"/pagess","entities":[{"offset":0,"length":6,"type":"bot_command"}]}}`

// Test client http request
func TestCheckNotFound(t *testing.T) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	w, err := http.Get("https://0.0.0.0:8080")

	if err != nil {
		panic(err)
	}

	if http.StatusNotFound != w.StatusCode {
		t.Errorf("Bad statusCode %v", w.StatusCode)
	}
}

// Test function handler
func TestCheckBadRequest(t *testing.T) {
	reader := strings.NewReader(readerJsonError)

	r, _ := http.NewRequest("POST", "/", reader)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v1.Index(w, r, httprouter.Params{})
	})

	handler.ServeHTTP(w, r)

	if http.StatusBadRequest != w.Code {
		t.Errorf("Bad statusCode %v", w.Code)
	}
}

// Test function handler
func TestCheckStatusOK(t *testing.T) {
	reader := strings.NewReader(readerJsonSuccess)

	r, _ := http.NewRequest("POST", "/", reader)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v1.Index(w, r, httprouter.Params{})
	})

	handler.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Bad statusCode %v", w.Code)
	}
}
