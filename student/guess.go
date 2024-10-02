package main

import "math"

// Function to calculate mean of a slice of  numbers
func calculateMean(numbers []float64) float64 {
	if len(numbers) == 0 {
        return 0.0
    }
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

// Function to calculate the standard deviation of a slice of numbers
func calculateStandardDeviation(numbers []float64, mean float64) float64 {
	sumOfSquares := 0.0
	for _, number := range numbers {
		sumOfSquares += math.Pow(number-mean, 2)
	}
	// Calculate Variance:
	variance := sumOfSquares / float64(len(numbers))

	return math.Sqrt(variance)
}
