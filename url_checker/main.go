package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	c := make(chan error)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	results := map[string]string{} // {}적어줌으로써 초기화해야 사용가능 or make 사용
	// var results = make(map[string]string)
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := "OK"
		if <-c != nil {
			result = "FAILED"
		}
		results[urls[i]] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, c chan error) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println("url:", url, "err:", err, "resp.StatusCode:", resp.StatusCode)
		c <- errRequestFailed
	}
	c <- nil
}
