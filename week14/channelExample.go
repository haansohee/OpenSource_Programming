// channel을 사용하여 동기화 문제를 해결하자~! (JAVA의 Synchronized)
// 우선, 스레드 간의 통신할 채널이 필요함. 동기화해야 하니깐..
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	size := make(chan int)
	go responseSize("http://www.naver.com", size)
	go responseSize("http://www.daum.net", size)
	go responseSize("http://www.nate.com", size)
	go responseSize("http://www.inhatc.ac.kr", size)
	// time.Sleep(5 * time.Second)
	fmt.Println(<-size)
	fmt.Println(<-size)
	fmt.Println(<-size)
	fmt.Println(<-size)

	fmt.Println("\n main 끝~!")
}

func responseSize(url string, channel chan int) {
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
	// fmt.Println(len(body))
	channel <- len(body) // 채널에 문자열의 길이 담기.
}
