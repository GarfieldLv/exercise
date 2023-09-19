package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if 3 == i {
				goto BREAKTAG
			}

			fmt.Printf("%d\t", j)
		}
	}

BREAKTAG: // label
	fmt.Println("over")

BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if 3 == i {
				break BREAKDEMO1
			}

			fmt.Printf("%d\t", j)
		}
	}

	fmt.Println("....")

CONTINUEDEMO:
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if 3 == i && 5 == j {
				continue CONTINUEDEMO
			}

			fmt.Printf("%d\t", j)
		}
	}

	fmt.Println("....")
}
