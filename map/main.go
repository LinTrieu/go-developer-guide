package main

import "fmt"

func main() {
    // approach 1
    // colors := map[string]string{
    //    "red": "#ff0000",
    //    "green": "#4bf745",
    // }

    // approach 2
    // var colors map[string]string

    // approach 3
    colors := make(map[string]string)

    // add key-value pair to an existing map
    colors["white"] = "#ffffff"
    fmt.Println(colors)

    // delete key-value pair to an existing map
    delete(colors, "white")
    fmt.Println(colors)
}
