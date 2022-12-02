package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	responseSize("http://www.naver.com")
	responseSize("http://www.daum.net")
	responseSize("http://www.nate.com")
	responseSize("http://www.inhatc.ac.kr")
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
	fmt.Println(len(body))
}
