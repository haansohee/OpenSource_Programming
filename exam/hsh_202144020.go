package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 팩토리얼 수로 변환하는 함수
// 팩토리얼 -> n!
// func Change(input int) int {
// 	var result int = 1
// 	var i int

// 	// fmt.Println("Change 로 넘어온 수 : ", input)

// 	for i = 1; i <= input; i++ { // 팩토리얼 계산
// 		result *= i
// 		// fmt.Println(result)
// 	}

// 	return result
// }

// 피보나치 수열로 변환 (f(n) = f(n-1) + f(n-2))
func Change(input int) int {
	// 재귀호출을 이용한 피보나치 수열
	// if input <= 2 {
	// 	return 1
	// } else {
	// 	return Change(input-1) + Change(input-2)
	// }

	var i int
	var last1 int
	var last2 int
	var result int

	if input <= 2 {
		return 1
	}

	last1 = 1
	last2 = 1

	for i = 2; i < input; i++ {
		result = last1 + last2
		last1 = last2
		last2 = result
	}

	return result
}

// 소수 여부 판단 함수  소수 : 1과 자신 외의 약수를 가지지 않는 1보다 큰 자연수
func Check(input int) bool {
	isPrime := true // 소수이면 true

	if input < 2 {
		isPrime := false
		return isPrime
	}

	// fmt.Println("Check 로 넘어온 수 : ", input)

	for i := 2; i < input; i++ {
		if input%i == 0 {
			isPrime = false
			return isPrime
		}
	}

	return isPrime
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	input, err := rd.ReadString('\n') // enter 칠 때까지 입력받고 input에 넣기

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input) // 불필요한 줄바꿈 제거

	inputSlice := strings.Split(input, " ")

	for _, num := range inputSlice {
		number, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(number)
		changeNum := Change(number)
		prime := Check(number)

		// if prime {
		// 	fmt.Printf("%d 팩토리얼 수 : %d / 소수여부 : true \n", number, changeNum)
		// } else {
		// 	fmt.Printf("%d 팩토리얼 수 : %d / 소수여부 : false \n", number, changeNum)
		// }

		if prime {
			fmt.Printf("%d 피보나치 수 : %d / 소수여부 : true \n", number, changeNum)
		} else {
			fmt.Printf("%d 피보나치 수 : %d / 소수여부 : false \n", number, changeNum)
		}
	}

}
