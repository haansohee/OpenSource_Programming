package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var quantity int
	// var l, w float64
	// l, w = 1.2, 2.4

	// var l, w float64 = 1.2, 2.4

	// var l, w = 1.2, 2.4

	// var l, w float32 = 1.2, 2.4

	// 단축 연산자 :=  -> 타입 및 var 선언 생략 가능.
	l, w := 1.2, 2.4
	quantity := 10

	var test bool
	fmt.Println(test)

	fmt.Printf("%.1f %.1f \n", l, w) // import해서 가져오는 클래스(클래스는 아니지만.. 쉽게 생각하기 위한...ㅎ)들은 모두 대문자로 시작함.
	// 대문자 : 외부 모듈에서 사용가능.(public), 소문자 : 모듈안에서만 접근 가능.
	fmt.Println(l*w, "제곱미터")

	// quantity = 10

	fmt.Println(quantity)
	fmt.Println(reflect.TypeOf(1))

	// fmt.Println(reflect.TypeOf(43))
	// fmt.Println(reflect.TypeOf(43.1))
	// fmt.Println(reflect.TypeOf(true))
	// fmt.Println(reflect.TypeOf("안녕")

	// go언어에서는 줄여쓰는 것을 권장. index -> i, maximum -> max

	// 형변환하는 방법은 다음과 같음.
	// var myInt int = 2
	// float64(myInt)
}
