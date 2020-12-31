package main

import "fmt"
import "os"

func main() {
	fmt.Println(os.Args[0])
	if len(os.Args) <= 1 {
		fmt.Println("usage")
		os.Exit(0)
	}
	fmt.Println(os.Args[1])
}