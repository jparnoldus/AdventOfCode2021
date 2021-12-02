package days

import (
	"strconv"
	"strings"
)

func Day2Challenge1(inputCh chan string, outputCh chan int) {
	translationsCh := make(chan Coordinates2D)
	go ConvertStringToTranslation(inputCh, translationsCh)

	positionsCh := make(chan Coordinates2D)
	go MovePosition(translationsCh, positionsCh)

	lastCh := make(chan Coordinates2D)
	go LastPosition(positionsCh, lastCh)
	go ConvertPositionToAnswer(lastCh, outputCh)
}

func Day2Challenge2(inputCh chan string, outputCh chan int) {
	translationsCh := make(chan Coordinates2D)
	go ConvertStringToTranslation(inputCh, translationsCh)

	positionsCh := make(chan Coordinates2D)
	go MovePositionWithAim(translationsCh, positionsCh)

	lastCh := make(chan Coordinates2D)
	go LastPosition(positionsCh, lastCh)
	go ConvertPositionToAnswer(lastCh, outputCh)
}

type Coordinates2D struct {
	X int
	Y int
}

func ConvertStringToTranslation(inputCh chan string, outputCh chan Coordinates2D) {
	defer close(outputCh)
	for translation := range inputCh {
		value, _ := strconv.Atoi(translation[len(translation)-1:])
		if strings.HasPrefix(translation, "forward") {
			outputCh <- Coordinates2D{X: value}
		} else if strings.HasPrefix(translation, "down") {
			outputCh <- Coordinates2D{Y: -value}
		} else if strings.HasPrefix(translation, "up") {
			outputCh <- Coordinates2D{Y: value}
		}
	}
}

func MovePosition(inputCh chan Coordinates2D, outputCh chan Coordinates2D) {
	defer close(outputCh)
	location := Coordinates2D{}
	for translation := range inputCh {
		location.X += translation.X
		location.Y += translation.Y
		outputCh <- location
	}
}

func MovePositionWithAim(inputCh chan Coordinates2D, outputCh chan Coordinates2D) {
	defer close(outputCh)
	aim := 0
	location := Coordinates2D{}
	for translation := range inputCh {
		aim -= translation.Y
		location.X += translation.X
		location.Y += translation.X * aim
		outputCh <- location
	}
}

func LastPosition(inputCh chan Coordinates2D, outputCh chan Coordinates2D) {
	defer close(outputCh)
	var last Coordinates2D
	for input := range inputCh {
		last = input
	}
	outputCh <- last
}

func ConvertPositionToAnswer(inputCh chan Coordinates2D, outputCh chan int) {
	defer close(outputCh)
	for input := range inputCh {
		outputCh <- input.X * input.Y
	}
}
