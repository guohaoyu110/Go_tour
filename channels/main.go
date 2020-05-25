// how we can use the concurrency things in go to fix our problems
// first build a website checker
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}
	c := make(chan string)
	// for _, link := range links {
	// 	// static tracing program
	// 	checkLink(link, c) // all checks are successful, all the website is up
	// }
	// create new channel

	// c is our channel , it's free to pass it to our
	for _, link := range links {
		go checkLink(link, c) // it's a go rountine
		// this is a child routine created by the go keyword
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep it's up!"
}

// integrate go rountines into our program
// the child go rountine doesn't show because go rountine ends when the main rountien ends
