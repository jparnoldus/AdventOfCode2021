package days

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay2Challenge1(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	inputCh := make(chan string)
	resultsCh := make(chan int)

	go Day2Challenge1(inputCh, resultsCh)

	for _, line := range input {
		inputCh <- line
	}
	close(inputCh)

	assert.Equal(t, -150, <-resultsCh)
}

func TestDay2Challenge2(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	inputCh := make(chan string)
	resultsCh := make(chan int)

	go Day2Challenge2(inputCh, resultsCh)

	for _, line := range input {
		inputCh <- line
	}
	close(inputCh)

	assert.Equal(t, 900, <-resultsCh)
}
