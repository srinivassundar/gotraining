package main

import "fmt"

func main() {

	fmt.Println("Im in the main function")
	foo()
	fmt.Println("I want something more")

	for i := 0; i < 50; i++ {
		fmt.Println(i)
	}
	bar()
}

func foo() {
	fmt.Println("Im in foo")

}

func bar() {
	fmt.Println("and then we exited")

}
