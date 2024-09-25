package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Read input numbers.

// Store recent numbers in a slice.
// Calculate moving average and standard deviation.
// Print range prediction.
const maxHistory = 5 // define how many nubers to keeo in history

func main() {
	var history []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		// output the current input
		fmt.Println(num)
		// add the new number to the history slice
		history = append(history, num)

		//Only predict when we have enough data
		if len(history) > 1 {
			// use only the last 'maxHistory' numbers for prediction

			startIndex := len(history) - maxHistory
			if startIndex < 0 {
                startIndex = 0
            }
			recentNumbers := history[startIndex:]

			lowerBound, upperBound := PredictRange(recentNumbers)
			fmt.Printf("Predicted range: [%d, %d]\n", lowerBound, upperBound)
		}
	}
}
