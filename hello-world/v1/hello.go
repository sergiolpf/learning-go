package main

import "fmt"

const englishHelloPrefis = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefis + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
