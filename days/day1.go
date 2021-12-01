package days

import (
	"strconv"
)

func Day1Challenge1(inputCh chan string, outputCh chan int) {
	intCh := make(chan int)
	go ConvertStringChannelToIntChannel(inputCh, intCh)

	greaterThanPreviousCh := make(chan int)
	go IsGreaterThanPrevious(intCh, greaterThanPreviousCh)
	go Count(greaterThanPreviousCh, outputCh)
}

func Day1Challenge2(inputCh chan string, outputCh chan int) {
	intCh := make(chan int)
	go ConvertStringChannelToIntChannel(inputCh, intCh)

	movingSumCh := make(chan int)
	go GetThreeDayMovingSum(intCh, movingSumCh)

	greaterThanPreviousCh := make(chan int)
	go IsGreaterThanPrevious(movingSumCh, greaterThanPreviousCh)
	go Count(greaterThanPreviousCh, outputCh)
}

func ConvertStringChannelToIntChannel(inputCh chan string, outputCh chan int) {
	defer close(outputCh)
	for line := range inputCh {
		depth, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		outputCh <- depth
	}
}

func IsGreaterThanPrevious(inputCh chan int, outputCh chan int) {
	defer close(outputCh)
	previous := 0
	for depth := range inputCh {
		if previous != 0 {
			if depth > previous {
				outputCh <- 1
			}
		}
		previous = depth
	}
}

func Count(inputCh chan int, outputCh chan int) {
	defer close(outputCh)
	counter := 0
	for _ = range inputCh {
		counter++
	}
	outputCh <- counter
}

func GetThreeDayMovingSum(inputCh chan int, outputCh chan int) {
	defer close(outputCh)
	var buffer []int
	for depth := range inputCh {
		buffer = append(buffer, depth)
		if len(buffer) > 2 {
			buffer = buffer[len(buffer)-3:]
			sum := 0
			for _, depth := range buffer {
				sum += depth
			}
			outputCh <- sum
		}
	}
}
