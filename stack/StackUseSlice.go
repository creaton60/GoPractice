package main

import (
	"strconv"
	"fmt"
	"strings"
)

/**
* 스택을 이용한 사칙연산과 괄호를 활용한 간단한 정수 게산기
 */
func Eval(expr string) int{
	var ops []string
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(higher string){
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op )< 0{
				//목록에 없는 연산자 이므로 종료
				return
			}

			ops = ops[:len(ops)-1]
			if op == "(" {
				//콸호를 제거 하였으므로 종료
				return
			}

			b,a := pop(), pop()
			switch op{
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}

	for _, token := range strings.Split(expr, " "){
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			//덧셈과 뺼셈과 동등한 우선순위를 적용
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			reduce("*/")
			ops = append(ops, token)
		case ")":
			reduce("+-*/(")
		default :
			num, _:= strconv.Atoi(token)
			nums = append(nums, num)


		}
	}
	reduce("+-*/")
	return nums[0]
}

func ExampleEval()  {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("1 + 2"))
	fmt.Println(Eval("1 - 2 + 3"))
	fmt.Println(Eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println(Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
}


func main() {
	ExampleEval()
}
