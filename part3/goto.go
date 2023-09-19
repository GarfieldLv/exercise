package main

func main() {
test:
	println("test.")

	for i := 0; i < 3; i++ {
		println(i)
	}

	goto test
}
