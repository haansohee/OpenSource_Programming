package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// 문자열을 convert(형변환)하는 패키지
)

func main() {
	// Project(3)
	fmt.Print("성적 입력 : ")
	reader := bufio.NewReader(os.Stdin) // 표준 입력
	// input := reader.ReadString('\n') // \n까지 입력받기
	// input, _ := reader.ReadString('\n') // 에러처리 option(1) : Underscore(_) 사용하면 error 무시 가능. 웬만하면 에러도 처리해 주어야함(Handler Error)
	input, err := reader.ReadString('\n')
	//log.Fatal(err) // option(2) : log패키지에 있는 Fatal 함수 사용하여 error log 출력. <nil> 이 출력되면 버그는 없었다는 뜻. 근데! log.Fatal() 사용하면 여기서 멈추니까 if문을 사용하기.
	// if "error" is not nil
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input) // 불필요한 줄바꿈, 캐리지 리턴 등의 이스케이프시퀀스를 제거해줌. 쉽게 말하자면 엔터를 제거해 준 것.

	score, err := strconv.ParseFloat(input, 64) // 엽입력된 문자열 input을 64float형으로 형변환.
	if err != nil {                             // Trim 안 할 경우, (2022/09/23 11:24:56 strconv.ParseFloat: parsing "100\n": invalid syntax) 에러 발생.
		log.Fatal(err)
	}

	var status string // 합격 불합격 여부를 저장할 변수 선언.

	if score >= 60 { // input 에러 발생 (mismatched types string and untyped int)
		status = "합격"
	} else {
		status = "불합격"
	}

	fmt.Println("점수 : ", score, " / ", status)

	// Project(2)
	// broken := "G# r#cks!"
	// replacer := strings.NewReplacer("#", "o")  // #을 o로 replace
	// fixed := replacer.Replace(broken)
	// fmt.Println(fixed)  // Go rocks!

	// Project(1)
	// 	// var now time.Time = time.Now()
	// 	now := time.Now()
	// 	// var year int = now.Year()
	// 	year := now.Year()

	// 	fmt.Println(year)  // 2022
	// 	fmt.Println(now.Month().String())  // September
}
