package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var URL = "https://api.stackexchange.com/search/advanced/?site=stackoverflow&pagesize=10"

func e(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type Item struct {
	Tags []string `json:"tags"`
	Title string `json:"title"`
	Answered bool `json:"is_answered"` 
	Link string `json:"link"`
}

type Items struct {
	Items []Item `json:"items"`
}

func Search(question string, page int) (items Items) {
	req, err := http.NewRequest("GET", URL, nil)	
	q := req.URL.Query()
	q.Add("title", question)
	q.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	dec.Decode(&items)
	e(err)
	return
}
