package days

import (
	"strconv"
)

type Day1 struct {
	answer1 int
	answer2 int
}

func (d *Day1) HandleInput(input chan string, results chan int) {
	challenge1 := make(chan int)
	challenge2 := make(chan []int)

	go d.handle1(challenge1)
	go d.handle2(challenge2)

	var buffer []int
	for line := range input {
		depth, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		challenge1 <- depth

		buffer = append(buffer, depth)
		if len(buffer) > 2 {
			buffer = buffer[len(buffer)-3:]
			challenge2 <- buffer
		}
	}

	close(challenge1)
	close(challenge2)

	results <- d.answer1
	results <- d.answer2
}

func (d *Day1) handle1(input chan int) {
	previous := 0
	for depth := range input {
		if previous != 0 {
			if depth > previous {
				d.answer1++
			}
		}
		previous = depth
	}
}

func (d *Day1) handle2(input chan []int) {
	previous := 0
	for depths := range input {
		sum := 0
		for _, depth := range depths {
			sum += depth
		}
		if previous != 0 {
			if sum > previous {
				d.answer2++
			}
		}
		previous = sum
	}
}
