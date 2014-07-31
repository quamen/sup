package auth

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
)

func Test_BasicAuth(t *testing.T) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte("foo:bar"))

	h := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello"))
	})
	m := negroni.New()
	m.Use(Basic("foo", "bar"))
	m.UseHandler(h)

	r, _ := http.NewRequest("GET", "foo", nil)
	recorder := httptest.NewRecorder()
	m.ServeHTTP(recorder, r)

	if recorder.Code != 401 {
		t.Error("Response not 401")
	}

	respBody := recorder.Body.String()
	if respBody == "hello" {
		t.Error("Auth block failed")
	}
	if respBody != "Not Authorized\n" {
	}

	recorder = httptest.NewRecorder()
	r.Header.Set("Authorization", auth)
	m.ServeHTTP(recorder, r)

	if recorder.Code == 401 {
		t.Error("Response is 401")
	}

	if recorder.Body.String() != "hello" {
		t.Error("Auth failed, got: ", recorder.Body.String())
	}
}
