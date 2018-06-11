package fortune

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var fortunes = []string{"大吉", "中吉", "小吉", "凶", "大凶"}

type Handler struct {
	Time time.Time
}

type fortune struct {
	Result string
	time   time.Time
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := fortune{time: h.Time}
	f.draw()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(f); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, buf.String())
}

func (f *fortune) draw() {
	if f.isBeginningOfTheYear() {
		f.Result = fortunes[0]
		return
	}
	f.Result = fortunes[rand.Intn(len(fortunes))]
}

func (f fortune) isBeginningOfTheYear() bool {
	if f.time.Month() == 1 && f.time.Day() <= 3 {
		return true
	}

	return false
}
