package main

import (
	"fmt"
	"os"
)

func main() {
	// go run chapter06_2.go inha computer
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])

	// 실행 결과
	// [/var/folders/mm/tvqg_2710hddmn5g0lvz4wj00000gn/T/go-build2098564400/b001/exe/chapter06_2 inha computer]
	// [inha computer]

}
