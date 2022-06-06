package main

import "fmt"

func SayHelloTo(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func main() {
	var message = SayHelloTo("pengu")

	fmt.Println(message)
}
