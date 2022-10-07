package main

import "fmt"

func main() {
	notes := [3]string{"do", "re", "mi"}

	// for i := 0; i <= 3; i++ { // panic: runtime error: index out of range [3] with length 3
	// for i := 0; i <= len(notes); i++ { // 얘도 문제..
	// 	fmt.Println(i, notes[i])
	// }

	// for i, note := range notes {  // 인덱스와 값 출력
	// 	fmt.Println(i, note)
	// }

	// for i := range notes {  // 인덱스만 출력
	// 	fmt.Println(i)
	// }

	for _, note := range notes { // 값만 출력
		fmt.Println(note)
	}
}
