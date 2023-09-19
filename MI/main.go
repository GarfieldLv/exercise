package main

import "fmt"

func main() {
	s, err := NewServer("127.0.0.1", 8888)
	if err != nil {
		fmt.Println("NewServer failed, err: ", err)
	}

	s.Start()
}
