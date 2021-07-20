package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

)

func TestCreat(t *testing.T) {
	// httpmock.Activate()
	// defer httpmock.DeactivateAndReset()

	// httpmock.RegisterResponder("GET", "/user/account",
	// 	httpmock.NewStringResponder(200, "hello"),
	// )
	// var resp *http.Response

	// resp, err := http.Get("/user/account")
	// if err != nil {
	// 	fmt.Println("error")
	// }

	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// assert.Equal(t, "hello", string(body))

	// This is another way to test API
	req, err := http.NewRequest("GET", "/user/account", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Create)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestMain(t *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(t.Run())
}
