package main

import (
	"advent-of-code-2021/days"
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make(chan string)
	results := make(chan int)
	go days.Day2Challenge2(input, results)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input <- scanner.Text()
	}
	close(input)

	log.Println(<-results)
}
