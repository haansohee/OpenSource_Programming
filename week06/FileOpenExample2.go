package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloat(fn string) ([3]float64, error) {
	var numbers [3]float64
	f, err := os.Open(fn)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return numbers, nil
}
func main() {
	prices, err := GetFloat("data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, price := range prices {
		fmt.Println(price)
	}
}
