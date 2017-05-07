package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	endpoint = "http://translate.naver.com/nmt.dic"
	key      = "7f099ab597bc397ef059d128e75dd5198c97d4938162a215cbe9ef00edde8e8c"
)

func main() {
	sourceLang := os.Args[1]
	targetLang := "ko"
	sourceText := os.Args[2]

	if sourceLang == "ko" {
		targetLang = "en"
	}

	data := url.Values{}
	data.Set("source", sourceLang)
	data.Add("target", targetLang)
	data.Add("text", sourceText)
	data.Add("passportKey", key)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", endpoint, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, _ := client.Do(r)
	defer res.Body.Close()

	if res.StatusCode < 400 {
		body, _ := ioutil.ReadAll(res.Body)
		data := map[string]map[string]map[string]interface{}{}
		_ = json.Unmarshal(body, &data)
		fmt.Println(data["message"]["result"]["translatedText"])
	}
}
