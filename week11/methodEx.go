// call by reference. 메서드의 매개변수로 주소를 전달... 그래서 두배 됨!!
package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n *= 2
}

func main() {
	number := Number(4)
	fmt.Println("원래 수 : ", number)
	number.Double()
	fmt.Println("두 배가 된 수 : ", number) // 그래서 두배가 안됨.
}

// call by value. 메서드의 receiver매개변수로 값을 전달받은 코드.
// package main

// import "fmt"

// type Number int

// func (n Number) Double() {
// 	n = n * 2
// }

// func main() {
// 	number := Number(4)
// 	fmt.Println("원래 수 : ", number)
// 	number.Double()
// 	fmt.Println("두 배가 된 수 : ", number) // 그래서 두배가 안됨.
// }
