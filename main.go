package main

import (
	"advent-of-code-2021/days"
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make(chan string)
	results := make(chan int)
	go days.Day3Challenge1(input, results)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input <- scanner.Text()
	}
	close(input)

	log.Println(<-results)
}
