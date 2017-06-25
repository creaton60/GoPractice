package main

import "fmt"

func main() {
	for i,  r := range "ABCDE"{
			fmt.Println(i, r)
	}
	fmt.Println(len("ABCDE"))

}
