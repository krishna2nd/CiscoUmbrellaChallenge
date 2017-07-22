package uhandlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestGETGlobalHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/urlinfo/1/google.com", nil)
	w := httptest.NewRecorder()
	GlobalHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result UploadResponse
	if err := json.Unmarshal(body, &result); err != nil {
		t.Error(err)
	}
	if result.Status != "safe" {
		t.Error("Expected result", string(body))
	}
}

func TestPOSTGlobalHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/urls", bytes.NewReader([]byte("gogle.com/q?wowwwww\ngogle.com/q?thisis-bad-url")))
	w := httptest.NewRecorder()
	GlobalHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result UploadResponse
	if err := json.Unmarshal(body, &result); err != nil {
		t.Error(err)
	}
	if result.Status != "updated" || result.Count != 2 {
		t.Error(`Expected upload result {"status":"updated","Count":2,"message":""}`)
	}
}
