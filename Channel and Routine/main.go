package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// for {
	for l := range c {
		// fmt.Println(<-c)
		// go checkLink(<-c, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Link is not working:", link)
		// c <- "Link is not working: " + link
		c <- link
		return
	}
	defer resp.Body.Close()
	fmt.Println("Link is working:", link)
	c <- link
}
