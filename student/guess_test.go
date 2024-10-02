package main

import (
	"math"
	"testing"
)

func TestCalculateMeanPositiveNumbers(t *testing.T) {
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expected := 3.0
	result := calculateMean(numbers)
	if result != expected {
		t.Errorf("calculateMean(%v) = %f; expected %f", numbers, result, expected)
	}
}

func TestCalculateMeanNegativeNumbers(t *testing.T) {
	numbers := []float64{-1.0, -2.0, -3.0, -4.0, -5.0}
	expected := -3.0
	result := calculateMean(numbers)
	if result != expected {
		t.Errorf("calculateMean(%v) = %f; expected %f", numbers, result, expected)
	}
}

func TestCalculateMeanEmptyArray(t *testing.T) {
	numbers := []float64{}
	expected := 0.0
	result := calculateMean(numbers)
	if result != expected {
		t.Errorf("calculateMean(%v) = %f; expected %f", numbers, result, expected)
	}
}

func TestCalculateMeanSingleNumber(t *testing.T) {
	numbers := []float64{42.0}
	expected := 42.0
	result := calculateMean(numbers)
	if result != expected {
		t.Errorf("calculateMean(%v) = %f; expected %f", numbers, result, expected)
	}
}

func TestCalculateMeanMixedNumbers(t *testing.T) {
	numbers := []float64{-2.0, 1.0, -3.0, 4.0, 5.0}
	expected := 1.0
	result := calculateMean(numbers)
	if result != expected {
		t.Errorf("calculateMean(%v) = %f; expected %f", numbers, result, expected)
	}
}

func TestCalculateMeanLargeNumbers(t *testing.T) {
	numbers := []float64{1e15, 2e15, 3e15, 4e15, 5e15}
	expected := 3e15
	result := calculateMean(numbers)
	if result != expected {
		t.Errorf("calculateMean(%v) = %e; expected %e", numbers, result, expected)
	}
}

func TestCalculateMeanSmallNumbers(t *testing.T) {
	numbers := []float64{1e-10, 2e-10, 3e-10, 4e-10, 5e-10}
	expected := 3e-10
	result := calculateMean(numbers)
	if math.Abs(result-expected) > 1e-15 {
		t.Errorf("calculateMean(%v) = %.15e; expected %.15e", numbers, result, expected)
	}
}

func TestCalculateMeanRepeatingNumbers(t *testing.T) {
	numbers := []float64{2.0, 2.0, 2.0, 2.0, 2.0}
	expected := 2.0
	result := calculateMean(numbers)
	if result != expected {
		t.Errorf("calculateMean(%v) = %f; expected %f", numbers, result, expected)
	}
}

func TestCalculateMeanDecimalNumbers(t *testing.T) {
	numbers := []float64{1.5, 2.7, 3.2, 4.9, 5.1}
	expected := 3.48
	result := calculateMean(numbers)
	if math.Abs(result-expected) > 1e-10 {
		t.Errorf("calculateMean(%v) = %.10f; expected %.10f", numbers, result, expected)
	}
}

func TestCalculateMeanWithNaNValues(t *testing.T) {
	numbers := []float64{1.0, math.NaN(), 3.0, math.NaN(), 5.0}
	result := calculateMean(numbers)
	if !math.IsNaN(result) {
		t.Errorf("calculateMean(%v) = %f; expected NaN", numbers, result)
	}
}
