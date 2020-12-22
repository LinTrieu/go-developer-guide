package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
