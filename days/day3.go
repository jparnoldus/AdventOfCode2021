package days

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

func Day3Challenge1(inputCh chan string, outputCh chan int) {
	splitCh := make(chan string)
	intCh := make(chan int)
	var distributorChs []chan int
	var answerChs []chan int
	for i := 0; i < 12; i++ {
		distributorChs = append(distributorChs, make(chan int))
		answerChs = append(answerChs, make(chan int))
	}
	combinedCh := make(chan []int)
	bitStringCh := make(chan string)
	duplicateCh := make(chan string)
	flippedCh := make(chan string)
	decimalsCh := make(chan int)

	go SplitStringIntoCharacters(inputCh, splitCh)
	go ConvertStringToInt(splitCh, intCh)
	go DistributeInt(intCh, distributorChs)
	for index, distributorCh := range distributorChs {
		go MostCommon(distributorCh, answerChs[index])
	}
	go CombineInOrder(answerChs, combinedCh)
	go CombineToBitString(combinedCh, bitStringCh)
	go DuplicateString(bitStringCh, duplicateCh)
	go FlipEveryOtherBitString(duplicateCh, flippedCh)
	go ConvertBitStringToInt(flippedCh, decimalsCh)
	go Multiply(decimalsCh, outputCh)
}

func Day3Challenge2(inputCh chan string, outputCh chan int) {
	inputChs := []chan string{inputCh}
	for i := 0; i < 1; i++ {
		filterCh := make(chan string)
		duplicateCh := make(chan string)
		// specificCh := make(chan string)
		// specificIntCh := make(chan int)
		// mostCommonCh := make(chan int)
		// filteredInputCh := make(chan string)
		debugCh1 := make(chan string)
		debugCh2 := make(chan string)

		go DistributeString(inputChs[i], []chan string{filterCh, duplicateCh})
		// go GetSpecificCharacters(filterCh, i, specificCh)
		// go ConvertStringToInt(specificCh, specificIntCh)
		// go MostCommon(specificIntCh, mostCommonCh)
		// go FilterWithAtIndex(duplicateCh, i, mostCommonCh, filteredInputCh)
		go Print(filterCh, debugCh1)
		go Print(duplicateCh, debugCh2)
		//inputChs = append(inputChs, filteredInputCh)
	}
}

func forever() {
	for {
		time.Sleep(time.Second)
	}
}

func Print(inputCh chan string, outputCh chan string) {
	defer close(outputCh)
	for input := range inputCh {
		outputCh <- input
	}
	panic("The end")
}

func FilterWithAtIndex(inputCh chan string, index int, valueCh chan int, outputCh chan string) {
	defer close(outputCh)
	value := strconv.Itoa(<-valueCh)
	for input := range inputCh {
		if strings.Split(input, "")[index] == value {
			outputCh <- input
		}
	}
}

func GetSpecificCharacters(inputCh chan string, index int, outputCh chan string) {
	defer close(outputCh)
	for input := range inputCh {
		outputCh <- strings.Split(input, "")[index]
	}

}

func DuplicateString(inputCh chan string, outputCh chan string) {
	defer close(outputCh)
	for input := range inputCh {
		outputCh <- input
		outputCh <- input
	}
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
		parsed, _ := strconv.ParseInt(input, 2, 64)
		outputCh <- int(parsed)
	}
}

func FlipEveryOtherBitString(inputCh chan string, outputCh chan string) {
	defer close(outputCh)
	flipflop := true
	for input := range inputCh {
		if flipflop {
			outputCh <- input
			flipflop = false
		} else {
			flipped := ""
			for _, character := range strings.Split(input, "") {
				if character == "0" {
					flipped += "1"
				} else {
					flipped += "0"
				}
			}
			outputCh <- flipped
			flipflop = true
		}
	}
}
