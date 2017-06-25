package main

import "fmt"

func Example_StrCat(){

	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s) //잘안쓰고
	fmt.Println(*ps) //요거를 더 많이 쓴데요

}

func main() {
	Example_StrCat()

}
