package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.inhatc.ac.kr")

	if err != nil {
		log.Fatal(err)
	}

	// url 주소가 잘못된 경우 err :
	// 	2022/12/02 10:29:35 Get "http://www.inhatc.ac.kr1": dial tcp: lookup www.inhatc.ac.kr1: no such host
	// exit status 1

	defer resp.Body.Close() // 메인 함수 종료 직전에 연결 해제. defer 키워드는 delay라고 생각하자. (코드를 다 받아올 때까지 지연시킨다는 것.)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	// resp.Body.Close() // 만약에 웹페이지 코드가 길거나 큰 규모의 페이지이고 서버가 느리다면 코드를 다 가져오기 전에 메인 함수가 바로 끝남.
}
