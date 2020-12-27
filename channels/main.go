package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// exercise 80
	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}

	// exercise 78
	// for l := range c {
	// time.Sleep(5 * time.Second)
	// go checkLink(l, c)
	// }

	// exercise 77
	// for {
	// 	go checkLink(<-c, c)
	// }

	// exercise 76
	// for i := 0; i < len(links); i++ {
	// fmt.Println(<-c)
	// }

	// exercise 75
	// fmt.Println(<-c) x repeat 5 times
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
