package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	. "github.com/stretchr/testify/require"
)

func TestCalculateTax_Under50k(t *testing.T) {
	expected := float32(10000.0)
	tax := calculateTax(10000)
	Equal(t, tax, expected)
}

func TestCalculateTax_Over50K(t *testing.T) {
	expected := float32(41500.0)
	tax := calculateTax(50000)
	Equal(t, tax, expected)
}

func TestCalculateTax_Over150K(t *testing.T) {
	expected := float32(109500.0)
	tax := calculateTax(150000)
	Equal(t, tax, expected)
}

func TestCalculateAverage_WithNegativeValue(t *testing.T) {
	expected := float32(3.5)
	avg, _ := calculateAverage(2.0, 3.0, 4.0, 5.0)
	Equal(t, avg, expected)
	_, err := calculateAverage(-1.0, -2.0)
	require.Errorf(t, err, "negative number: -1.00")
}

func TestCalculateAverage_HappyPath(t *testing.T) {
	expected := float32(3.5)
	avg, _ := calculateAverage(2.0, 3.0, 4.0, 5.0)
	Equal(t, avg, expected)
}

func TestCalculateSalary_CategoryA(t *testing.T) {
	expected := float32(45000.0)
	salary := calculateSalary(10, "A")
	Equal(t, salary, expected)
}

func TestCalculateSalary_CategoryB(t *testing.T) {
	expected := float32(18000.0)
	salary := calculateSalary(10, "B")
	Equal(t, salary, expected)
}

func TestCalculateSalary_CategoryC(t *testing.T) {
	expected := float32(10000.0)
	salary := calculateSalary(10, "C")
	Equal(t, salary, expected)
}

func TestCalculateSalary_DefaultCase(t *testing.T) {
	expected := float32(0.0)
	salary := calculateSalary(10, "D")
	Equal(t, salary, expected)
}

func TestOperation_InvalidOperation(t *testing.T) {
	_, err := operation("invalid")
	Error(t, err) // Check if an error is returned.
	//require.Errorf(t, err, "invalid operation type")
}

func TestMinFunc_HappyPath(t *testing.T) {
	expected := 2
	minFunc, _ := operation(minimum)
	minValue := minFunc(3, 2, 3, 4, 10, 2, 4, 5)
	Equal(t, minValue, expected)
}

func TestAverageFunc_HappyPath(t *testing.T) {
	expected := 3
	averageFunc, _ := operation(average)
	averageValue := averageFunc(3, 2, 3, 4, 1, 2, 4, 5)
	Equal(t, averageValue, expected)
}

func TestMaxFunc_HappyPath(t *testing.T) {
	expected := 5
	maxFunc, _ := operation(maximum)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	Equal(t, maxValue, expected)
}

func TestAnimalFunc_DogHappyPath(t *testing.T) {
	const (
		dog = "dog"
	)

	animalDogFunc, _ := animal(dog)
	value := animalDogFunc(10)
	Equal(t, value, float64(100))
}

func TestAnimalFunc_CatHappyPath(t *testing.T) {
	const (
		cat = "cat"
	)

	animalCatFunc, _ := animal(cat)
	value := animalCatFunc(10)
	Equal(t, value, float64(50))
}

func TestAnimalFunc_HamsterHappyPath(t *testing.T) {
	const (
		hamster = "hamster"
	)

	animalHamsterFunc, _ := animal(hamster)
	value := animalHamsterFunc(10)
	Equal(t, value, float64(2.5))
}

func TestAnimalFunc_TarantulaHappyPath(t *testing.T) {
	const (
		tarantula = "tarantula"
	)

	animalTarantulaFunc, _ := animal(tarantula)
	value := animalTarantulaFunc(10)
	Equal(t, value, float64(1.5))
}
