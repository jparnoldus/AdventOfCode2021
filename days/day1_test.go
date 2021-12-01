package days

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay1_HandleInput(t *testing.T) {
	input := []string{"199","200","208","210","200","207","240","269","260","263"}
	inputCh := make(chan string)
	resultsCh := make(chan int)

	d := Day1{}
	go d.HandleInput(inputCh, resultsCh)

	for _, line := range input {
		inputCh <- line
	}
	close(inputCh)

	assert.Equal(t, 7, <-resultsCh)
	assert.Equal(t, 5, <-resultsCh)
}