package main

import "fmt"

func main() {
	c, err := InitClarum(Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
