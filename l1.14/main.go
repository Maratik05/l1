package main

import "fmt"

func vType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool, chan interface{}:
		return "chan"
	default:
		return "..."
	}
}

func main() {
	a := make(chan int)
	fmt.Print(vType(a))

}
