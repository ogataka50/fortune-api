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
}

type drawing struct {
	time time.Time
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := drawing{time: h.Time}
	f := fortune{Result: d.draw()}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(f); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, buf.String())
}

func (d drawing) draw() string {
	if d.isBeginningOfTheYear() {
		return fortunes[0]
	}
	return fortunes[rand.Intn(len(fortunes))]
}

func (d drawing) isBeginningOfTheYear() bool {
	if d.time.Month() == 1 && d.time.Day() <= 3 {
		return true
	}

	return false
}
