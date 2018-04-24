package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"os"
)

const (
	baseUrl = "https://sitereview.bluecoat.com/resource/lookup"
	userAgent = "Mozilla/5.0"
	contentType = "application/json"
)

func checkUrl(url string) {
	var jsonStr= []byte(fmt.Sprintf(`{"url":"%s","captcha":""}`, url))
	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	checkUrl(os.Args[1])
}
