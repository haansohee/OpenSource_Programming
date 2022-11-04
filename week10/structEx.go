package main

import (
	"fmt"
	"magazine"
)

func main() {
	var s magazine.Subscriber // 구독자
	var e magazine.Employee   // 직원

	s.Name = "한소희"
	s.Rate = 5.99
	s.Active = true

	e.Name = "김직원"
	// e.HomeAddress.City = "인천"  // 이걸 줄여서 써보자!
	e.City = "인천"

	fmt.Println(s.Name, e.Name)
	fmt.Println(e.City)
}
