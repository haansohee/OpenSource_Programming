package main

// 랜덤으로 생성된 값을 추측하는 게임 프로그램.

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// rand.Seed(99)
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("1에서 100사이의 수를 추측하세요.")

	// 입력받기
	reader := bufio.NewReader(os.Stdin) // 키보드 입력받기
	success := false                    // 추측 성공 여부 변수

	for guesses := 0; guesses < 10; guesses++ { // 10번의 기회
		fmt.Println("남은 기회 : ", (10 - guesses), "회")
		fmt.Print("추착한 값 입력 : ")

		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input) // string을 int로 변환

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("정답은 추측하신 값보다 더 큰 값입니다.")
		} else if guess > target {
			fmt.Println("정답은 추측하신 값보다 더 작은 값입니다.")
		} else {
			success = true // 추측 성공
			fmt.Println("정답입니다.")
			break // for문 탈출
		}
	}

	if !success {
		fmt.Println("주어진 기회를 다 사용하셨습니다. 정답은 ", target, "입니다.")
	}

}
