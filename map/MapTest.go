package main

import "fmt"

func main() {

	var a map[string]int = make(map[string]int) //make 함수로 키는 string 값은 int 인 맵

	var b = make(map[string]string) //맵을 선언할 때 map 키워드와 자료형 생략 가능
	c := make(map[string]int) //덕타이핑으로 var 도 생략 가능


	a["Hello"] = 10
	b["world"] = "Test"
	c["Hello"] = 30

	fmt.Println(a["Hello"]) //10
	fmt.Println(b["Hello"]) //?
	fmt.Println(c["Hello"]) //30

}
