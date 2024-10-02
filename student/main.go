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
const maxHistory = 10 // define how many nubers to keeo in history

func main() {
	var history []float64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		// output the current input
		fmt.Println(num)
		// add the new number to the history slice
		history = append(history, num)
		//
		if len(history) > maxHistory {
			history = history[len(history)-maxHistory:]
		}
		// Only predict when we have enough data
		if len(history) > 1 {
			// calculate the mean and standard deviation
			mean := calculateMean(history)
			stdDev := calculateStandardDeviation(history, mean)

			// predict the next range using the mean and standard deviation
			lowerBound := int(mean - 2*stdDev)
			upperBound := int(mean + 2*stdDev)
			fmt.Printf("Predicted range: [%d, %d]\n", lowerBound, upperBound)
		}
	}
}
