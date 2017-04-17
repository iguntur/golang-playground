package main;

import "fmt";

func sayHello(input string) string {
	return "Hello, " + input + "!";
}

func main() {
	fmt.Println(sayHello("go"));
}
