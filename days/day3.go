package days

import (
	"strconv"
	"strings"
	"sync"
)

func Day3Challenge1(inputCh chan string, outputCh chan int) {
	splitCh := make(chan string)
	go SplitStringIntoCharacters(inputCh, splitCh)
	intCh := make(chan int)
	go ConvertStringToInt(splitCh, intCh)

	distributorChs := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}
	go DistributeInt(intCh, distributorChs)

	var answerChs []chan int
	for _, distributorCh := range distributorChs {
		answerCh := make(chan int)
		go MostCommon(distributorCh, answerCh)
		answerChs = append(answerChs, answerCh)
	}

	combinedCh := make(chan []int)
	go CombineInOrder(answerChs, combinedCh)

	bitStringCh := make(chan string)
	go CombineToBitString(combinedCh, bitStringCh)

	bitStringChs := []chan string{
		make(chan string),
		make(chan string),
	}
	go DistributeString(bitStringCh, bitStringChs)

	flippedCh := make(chan string)
	go FlipBitString(bitStringChs[1], flippedCh)

	bitStringChs = append(bitStringChs[:1], flippedCh)
	combinedStringCh := make(chan string)
	go Combine(bitStringChs, combinedStringCh)

	decimalsCh := make(chan int)
	go ConvertBitStringToInt(combinedStringCh, decimalsCh)

	go Multiply(decimalsCh, outputCh)
}

func Multiply(inputCh chan int, outputCh chan int) {
	defer close(outputCh)
	result := 1
	for input := range inputCh {
		result *= input
	}
	outputCh <- result
}

func CombineInOrder(inputChs []chan int, outputCh chan []int) {
	defer close(outputCh)
	wg := &sync.WaitGroup{}
	wg.Add(len(inputChs))
	for index, inputCh := range inputChs {
		go func(index int, inputCh chan int, wg *sync.WaitGroup) {
			for input := range inputCh {
				outputCh <- []int{index, input}
				wg.Done()
			}
		}(index, inputCh, wg)
	}
	wg.Wait()
}

func Combine(inputChs []chan string, outputCh chan string) {
	defer close(outputCh)
	wg := &sync.WaitGroup{}
	wg.Add(len(inputChs))
	for _, inputCh := range inputChs {
		go func(inputCh chan string, wg *sync.WaitGroup) {
			for input := range inputCh {
				outputCh <- input
			}
			wg.Done()
		}(inputCh, wg)
	}
	wg.Wait()
}

func DistributeInt(inputCh chan int, outputChs []chan int) {
	counter := 0
	for input := range inputCh {
		outputChs[counter] <- input
		counter++
		if counter == len(outputChs) {
			counter = 0
		}
	}
	for _, outputCh := range outputChs {
		close(outputCh)
	}
}

func DistributeString(inputCh chan string, outputChs []chan string) {
	counter := 0
	for input := range inputCh {
		outputChs[counter] <- input
		counter++
		if counter == len(outputChs) {
			counter = 0
		}
	}
	for _, outputCh := range outputChs {
		close(outputCh)
	}
}

func CombineToBitString(inputCh chan []int, outputCh chan string) {
	defer close(outputCh)
	bitStringMap := make(map[int]int)
	for bit := range inputCh {
		bitStringMap[bit[0]] = bit[1]
	}
	bitString := ""
	for i := 0; i < len(bitStringMap); i++ {
		bitString += strconv.Itoa(bitStringMap[i])
	}
	outputCh <- bitString
}

func SplitStringIntoCharacters(inputCh chan string, outputCh chan string) {
	defer close(outputCh)
	for input := range inputCh {
		for _, character := range strings.Split(input, "") {
			outputCh <- character
		}
	}
}

func MostCommon(inputCh chan int, outputCh chan int) {
	defer close(outputCh)
	occurrences := make(map[int]int)
	for input := range inputCh {
		occurrences[input]++
	}
	mostCommon := []int{0, 0}
	for key, value := range occurrences {
		if value > mostCommon[1] {
			mostCommon = []int{key, value}
		}
	}
	outputCh <- mostCommon[0]
}

func ConvertBitStringToInt(inputCh chan string, outputCh chan int) {
	defer close(outputCh)
	for input := range inputCh {
		parsed, _ := strconv.ParseInt(input, 2, 5)
		outputCh <- int(parsed)
	}
}

func FlipBitString(inputCh chan string, outputCh chan string) {
	defer close(outputCh)
	for input := range inputCh {
		flipped := ""
		for _, character := range strings.Split(input, "") {
			if character == "0" {
				flipped += "1"
			} else {
				flipped += "0"
			}
		}
		outputCh <- flipped
	}
}
