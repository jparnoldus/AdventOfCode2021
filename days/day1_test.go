package days

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Challenge1(t *testing.T) {
	input := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	inputCh := make(chan string)
	resultsCh := make(chan int)

	go Day1Challenge1(inputCh, resultsCh)

	for _, line := range input {
		inputCh <- line
	}
	close(inputCh)

	assert.Equal(t, 7, <-resultsCh)
}

func TestDay1Challenge2(t *testing.T) {
	input := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	inputCh := make(chan string)
	resultsCh := make(chan int)

	go Day1Challenge2(inputCh, resultsCh)

	for _, line := range input {
		inputCh <- line
	}
	close(inputCh)

	assert.Equal(t, 5, <-resultsCh)
}
