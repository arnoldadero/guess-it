package main

import "math"

// Function to calculate mean of a slice of  numbers
func calculateMean(numbers []int) float64 {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
}

// Function to calculate the standard deviation of a slice of numbers
func calculateStandardDeviation(numbers []int, mean float64) float64 {
	sumOfSquares := 0.0
	for _, number := range numbers {
		sumOfSquares += math.Pow(float64(number)-mean, 2)
	}
	return math.Sqrt(sumOfSquares / float64(len(numbers)))
}

// Function to predict the next range of numbers based on the given history
func PredictRange(numbers []int) (int, int) {
	mean := calculateMean(numbers)
	standardDeviation := calculateStandardDeviation(numbers, mean)

	lowerBound := int(mean - standardDeviation)
	upperBound := int(mean + standardDeviation)

	// Ensure the range is always non-negative
	if lowerBound < 0 {
		lowerBound = 0
	}

	return lowerBound, upperBound
}
