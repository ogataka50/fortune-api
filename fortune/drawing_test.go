package fortune

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	h := Handler{Time: time.Now()}
	h.ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	_, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestHandlerNewYear(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	h := Handler{Time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)}
	h.ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	var f fortune
	if err := json.NewDecoder(rw.Body).Decode(&f); err != nil {
		t.Fatal("unexpected error")
	}

	if f.Result != "大吉" {
		t.Fatal("unexpected fortune...")
	}
}
