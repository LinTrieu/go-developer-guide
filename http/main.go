package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(resp)
}
