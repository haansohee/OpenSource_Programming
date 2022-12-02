package main

import (
	"fmt"
	"math"
)

// 직육면체와 구를 통칭할 수 있는 인터페이스를 만들어 보자. "도형"이라는 인터페이스를 만들면 되겠죠?
type Shape3D interface { // 3차원 도형이라는 인터페이스
	Volume() float64 // 부피와 겉넓이를 구하는 메서드.
	Surface() float64
	// 덕 타이핑!
}

type Cuboid struct { // 직육면체. 직육면체의 필드 : 너비, 높이, 길이
	width  float64
	height float64
	length float64
}

// 부피와 겉넓이를 구하는 메서드

func (c Cuboid) Volume() float64 { // 부피
	return c.width * c.height * c.height
}

func (c Cuboid) Surface() float64 {
	area := c.length * c.width * 2 // 여섯개의 면 중 하나만 구하고 반대편의 면까지 해서 (*2)한 것. 4개의 면 더 구하면 됨.
	area = area + (c.length * c.height * 2)
	area = area + (c.width * c.height * 2)

	return area
}

type Sphere struct { // 구. 구의 필드 : 반지름
	radius float64
}

func (s Sphere) Volume() float64 { // 부피 공식 : 4/3πr³
	// return (4.0 / 3.0) * math.Pi * s.radius * s.radius * s.radius
	return (4.0 / 3.0) * math.Pi * math.Pow(s.radius, 3.0)
}

func (s Sphere) Surface() float64 { // 겉넓이 공식 : 4πr2
	return 4.0 * math.Pi * math.Pow(s.radius, 2.0)
}

//	func PrintInfo(c Cuboid) {
//		fmt.Println(c)
//		fmt.Printf("부피 : %0.2f \n", c.Volume())
//		fmt.Printf("겉넓이 : %0.2f \n", c.Surface())
//	}
func PrintInfo(shape3D Shape3D) {
	fmt.Println(shape3D)
	fmt.Printf("부피 : %0.2f \n", shape3D.Volume())
	fmt.Printf("겉넓이 : %0.2f \n", shape3D.Surface())
}

func main() {
	c1 := Cuboid{width: 5.0, height: 10.5, length: 2.5}
	PrintInfo(c1)

	s1 := Surface{radius: 10.0}
	PrintInfo(s1)
}
