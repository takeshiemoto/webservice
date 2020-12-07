package main

import "fmt"

func main() {
	post := Post{
		ID:      1,
		Content: "Hello",
		Author:  "World",
	}
	fmt.Printf("%+v\n", post)
}
