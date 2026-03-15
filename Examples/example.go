package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	request, _ := http.NewRequest("POST", "http://localhost:42690/go/pher", nil)

	request.Host = "cf.erisa.uk"

	headers := map[string]string{
		"x-tls-url":            "https://cf.erisa.uk/", // target
		"x-tls-profile":        "firefox_147", // supports latest tls-client profiles
		"x-tls-method": 		"GET", // here you can specify the method 
		"x-tls-proxy":          "", // you can insert in this format ip:port:username:password, leave it empty if u wanna use it without a proxy
		"x-tls-timeout":        "10",
		"x-tls-force-h1": 		"false", // optional
		"x-tls-follow-redirects": 	"true", // optional, true or false
		"x-tls-insecure-skip-verify": 	"false", // optional too
		"x-tls-header-order":   "Accept,User-Agent,Content-Type",
		"x-tls-pseudo-order":   ":method,:authority,:path,:scheme", // h2 order
		"User-Agent":           "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:147.0) Gecko/20100101 Firefox/147.0",
		"Accept":               "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Encoding":      "gzip, deflate, br, zstd",
		"Connection":           "keep-alive",
		"Sec-Fetch-Mode":       "navigate",
		"Sec-Fetch-Site":       "none",
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}