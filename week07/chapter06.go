package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloat(fn string) ([]float64, error) {
	// var numbers [3]   // 배열의 크기는 3인데 txt파일에는 3 이상의 요소들이 들어 있다면 > panic: runtime error: index out of range [3] with length 3 발생 그럼!! slice 사용하자.
	var numbers []float64

	f, err := os.Open(fn)
	if err != nil {
		return numbers, err
	}
	// i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		// i++
		numbers = append(numbers, number)
	}
	err = f.Close()
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	if scanner.Err() != nil {
		// log.Fatal(scanner.Err())
		return nil, scanner.Err()
	}
	return numbers, nil
}
func main() {
	numbers, err := GetFloat("data2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		// fmt.Println(price)
		fmt.Println(number)
		sum += number
	}

	fmt.Printf("합계 : %0.2f \n", sum)
	fmt.Printf("평균 : %0.2f \n", sum/(float64(len(numbers))))
}
