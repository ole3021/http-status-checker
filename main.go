package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://changbaispring.com",
		"http://ole3021.me",
		"http://sspai.com",
		"http://www.google.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkStatus(url, c)
	}

	for url := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkStatus(url, c)
		}(url)
	}
}

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}
	fmt.Println(url, "is ok.")
	c <- url
}
