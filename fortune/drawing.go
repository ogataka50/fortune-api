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

type fortune struct {
	Result string
}

func draw() string {
	return fortunes[rand.Intn(len(fortunes))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func Handler(w http.ResponseWriter, r *http.Request) {
	f := fortune{
		Result: draw(),
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(f); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, buf.String())
}
