// go 키워드는 Thread라고 생각하면 됨.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSize("http://www.naver.com")
	go responseSize("http://www.daum.net")
	go responseSize("http://www.nate.com")
	go responseSize("http://www.inhatc.ac.kr")
	time.Sleep(5 * time.Second) // 5초 delay (위 스레드들이 다 작동하기 전에 메인이 먼저 끝날 수 있으니까 delay시켜줌.) 하지만.. 이상적인 방법 아님. -> channel을 사용하자!

	fmt.Println("\n main 끝~!")
}

func responseSize(url string) {
	fmt.Println("Getting", url)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))
	fmt.Println(len(body)) // 다섯개의 스레드(메인 함수까지)가 동시에 실행되니까 출력되는 결과가 내가 원하는 순서대로 나오지 않음. (네이버 -> 다음 -> 네이트 -> 학교공홈)
}
