/*
학번 : 202144020
성명 : 한소희

13주차 수업때 사용했던 코드를 수정하여 병렬처리 프로그램을 완성하시오.
조건1)
원기둥 구조체와 관련 메서드를 작성하시오.
조건2)
main함수 Shape3D타입의 슬라이스 변수 shapes를 선언하고
직육면체 변수 1개, 원기둥 변수 2개, 구 변수 1개를 담으시오.
조건3)
PrintInfo함수 안쪽에 타입 단언(type assertion)을 사용하시오.
조건4)
고루틴을 사용하시오. (채널은 사용하지 않아도 됨)

// 원기둥의 부피 및 겉넓이
// V = πr2 * h
// A = 2πr2 + 2πrh

다음은 프로그램 출력결과

PS D:\gowork\week15_01> go run .\final01.go
{10 10} 원기둥의 부피는 3141.59 입니다
겉넓이는 1256.64 입니다
{10 2 5} 직육면체의 부피는 100.00 입니다
겉넓이는 160.00 입니다
{10} 구의 부피는 4188.79 입니다
겉넓이는 1256.64 입니다
{6 2} 원기둥의 부피는 75.40 입니다
겉넓이는 100.53 입니다
PS D:\gowork\week15_01> go run .\final01.go
{6 2} 원기둥의 부피는 75.40 입니다
겉넓이는 100.53 입니다
{10} 구의 부피는 4188.79 입니다
겉넓이는 1256.64 입니다
{10 2 5} 직육면체의 부피는 100.00 입니다
겉넓이는 160.00 입니다
{10 10} 원기둥의 부피는 3141.59 입니다
겉넓이는 1256.64 입니다
*/

package main

import (
	"fmt"
	"math"
	"time"
)

type Shape3D interface { // 3차원 도형이라는 인터페이스
	Volume() float64 // 부피와 겉넓이를 구하는 메서드.
	Surface() float64
	// 덕 타이핑!
}

// 조건1) 원기둥 구조체와 관련 메서드 작성
type Silnder struct { // 조건1-1) 구조체
	// 반지름과 높이 필드
	radius float64
	height float64
}

func (s Silnder) Volume() float64 { // 부피
	return math.Pi * math.Pow(s.radius, 2.0) * s.height
}

func (s Silnder) Surface() float64 { // 겉넓이
	// V = πr2 * h
	// A = 2πr2 + 2πrh
	return 2*math.Pi*math.Pow(s.radius, 2.0) + 2*math.Pi*s.radius*s.height
}

type Cuboid struct { // 직육면체. 직육면체의 필드 : 너비, 높이, 길이
	width  float64
	height float64
	length float64
}

func (c Cuboid) Volume() float64 {
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
	return (4.0 / 3.0) * math.Pi * math.Pow(s.radius, 3.0)
}

func (s Sphere) Surface() float64 { // 겉넓이 공식 : 4πr2
	return 4.0 * math.Pi * math.Pow(s.radius, 2.0)
}

func PrintInfo(shape3D Shape3D) {
	// 조건 3) PrintInfo함수 안쪽에 타입 단언(type assertion)을 사용 -> ..??

	// c, ok := shape3D.(Cuboid)

	if c, ok := shape3D.(Cuboid); ok {
		fmt.Print(c)
		fmt.Printf("정육면체 부피는 %0.2f 입니다. \n", shape3D.Volume())
		fmt.Printf("겉넓이는 %0.2f 입니다. \n", shape3D.Surface())
		fmt.Println()
	}

	if sp, ok := shape3D.(Sphere); ok {
		fmt.Print(sp)
		fmt.Printf("구 부피는 %0.2f 입니다. \n", shape3D.Volume())
		fmt.Printf("겉넓이는 %0.2f 입니다. \n", shape3D.Surface())
		fmt.Println()
	}

	if si, ok := shape3D.(Silnder); ok {
		fmt.Print(si)
		fmt.Printf("원기둥 부피는 %0.2f 입니다. \n", shape3D.Volume())
		fmt.Printf("겉넓이는 %0.2f 입니다. \n", shape3D.Surface())
		fmt.Println()
	}

	// if name, ok := shape3D.(Sphere) {
	// 	fmt.Print(shape3D)
	// 	fmt.Println()
	// 	fmt.Printf("%부피는 %0.2f 입니다. \n", name, shape3D.Volume())
	// 	fmt.Printf("%겉넓이는 %0.2f 입니다. \n", name, shape3D.Surface())
	// }

	// if name, ok := shape3D.(Silnder) {
	// 	fmt.Print(shape3D)
	// 	fmt.Println()
	// 	fmt.Printf("%부피는 %0.2f 입니다. \n", name, shape3D.Volume())
	// 	fmt.Printf("%겉넓이는 %0.2f 입니다. \n", name, shape3D.Surface())
	// }

}

func main() {
	// 조건2) main함수 Shape3D타입의 슬라이스 변수 shapes를 선언하고
	// 직육면체 변수 1개, 원기둥 변수 2개, 구 변수 1개를 담으시오.

	// c1 := Cuboid{width: 5.0, height: 10.5, length: 2.5}
	// PrintInfo(c1)

	// s1 := Surface{radius: 10.0}
	// PrintInfo(s1)
	shapes := make([]Shape3D, 4)

	shapes[0] = Cuboid{width: 10.0, height: 2.0, length: 5.0} // 직육면체
	shapes[1] = Sphere{radius: 10.0}                          // 구
	shapes[2] = Silnder{radius: 6.0, height: 2.0}             // 원기둥 1
	shapes[3] = Silnder{radius: 10.0, height: 10.0}           // 원기둥 2

	for i, _ := range shapes {
		// 조건 4) 고루틴 사용
		go PrintInfo(shapes[i])
	}

	time.Sleep(time.Second * 5)
}
