package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	pages := make(chan Page)
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
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("주소: %s 길이: %d \n", page.URL, page.Size)
	}

	fmt.Println("\n main 끝~!")
}

func responseSize(url string, channel chan Page) {
	// fmt.Println("Getting", url)

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
	// channel <- len(body) // 채널에 문자열의 길이 담기.
	channel <- Page{URL: url, Size: len(body)}
}
