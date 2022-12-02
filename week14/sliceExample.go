package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	sizes := make(chan int)
	urls := []string{
		"http://www.naver.com",
		"http://www.daum.net",
		"http://www.nate.com",
		"http://www.inhatc.ac.kr",
	}
	// go responseSize("http://www.naver.com", size)
	// go responseSize("http://www.daum.net", size)
	// go responseSize("http://www.nate.com", size)
	// go responseSize("http://www.inhatc.ac.kr", size)
	// time.Sleep(5 * time.Second)
	// fmt.Println(<-size)
	// fmt.Println(<-size)
	// fmt.Println(<-size)
	// fmt.Println(<-size)

	for _, url := range urls {
		go responseSize(url, sizes)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}

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
