package main // Defines the package name. This file belongs to the main package.

import "fmt" // Imports the fmt package for formatted I/O operations.

const (
	minimum = "minimum"

	average = "average"

	maximum = "maximum"
)

func main() {
	// This is the entry point of the program.
	fmt.Println("Hello, World!") // Prints "Hello, World!" to the console.

	executeExcercise("calculator")

	animalDog, msg := animal(dog)

	fmt.Println(animalDog(10), msg)

	animalCat, msg := animal(cat)

	fmt.Println(animalCat(10), msg)

	var amount float64

	amount += animalDog(10)

	fmt.Println(amount, msg)

	amount += animalCat(10)

	fmt.Println(amount, msg)
}

func executeExcercise(operationType string) {
	// Currently, this function is empty and does nothing.
	switch operationType {
	case "calculator":
		minFunc, _ := operation(minimum)

		averageFunc, _ := operation(average)

		maxFunc, _ := operation(maximum)

		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)

		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println("Minimum value:", minValue)
		fmt.Println("Average value:", averageValue)
		fmt.Println("Maximum value:", maxValue)

	case "tax":
		salary := 10000.0
		tax := calculateTax(float32(salary))
		fmt.Printf("Salary: %.2f, Tax: %.2f\n", salary, tax)

	case "average":
		numbers := []float32{2.5, 3.5, 4.5, 5.5}
		avg, _ := calculateAverage(numbers...)
		fmt.Printf("Average: %.2f\n", avg)
	case "salary":
		hours := 40
		category := "A"
		salary := calculateSalary(hours, category)
		fmt.Printf("Salary: %.2f\n", salary)
	default:
		fmt.Println("Invalid operation type")
	}
}

func calculateTax(salary float32) float32 {

	if salary < 50000 {
		return salary
	}
	// Calculates tax based on the salary.
	if salary >= 50000 && salary < 150000 {
		// If salary is not between 50,000 and 150,000,
		// subtracts 17% of the salary as tax.
		return salary - (salary * 0.17)
	}
	// For salaries between 50,000 and 150,000,
	// subtracts 27% of the salary as tax.
	return salary - (salary * 0.27)
}

func calculateAverage(numbers ...float32) (float32, error) {
	// Calculates the average of a variable number of float32 numbers.
	var total float32 // Initializes the total sum of numbers to 0.
	for _, number := range numbers {
		// Iterates over each number in the input slice.
		if number < 0 {
			return 0, fmt.Errorf("negative number: %.2f", number)
		}
		total += number // Adds the current number to the total sum.
	}
	// Divides the total sum by the number of elements to find the average.
	// Converts the length of the numbers slice to float32 for division.
	return total / float32(len(numbers)), nil // Returns the average value.
}

func calculateSalary(hours int, category string) float32 {
	// Calculates the salary based on the number of hours worked and the category.
	var totalSalary float32 // Initializes the total salary to be returned.
	switch category {
	case "A":
		totalSalary = float32(hours) * 3000
		return totalSalary + (totalSalary * 0.50)
	case "B":
		totalSalary = float32(hours) * 1500
		return totalSalary + (totalSalary * 0.20)
	case "C":
		return float32(hours) * 1000 // Sets the hourly rate to 1000 for category C.
	default:
		return 0
	}
}

func operation(operationType string) (func(...int) int, error) {
	// Returns the appropriate function based on the operation type.
	switch operationType {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		return nil, fmt.Errorf("unsupported operation: %s", operationType)
	}
}

func minFunc(numbers ...int) int {
	// Finds the minimum value among the input numbers.
	min := numbers[0] // Initializes the minimum value with the first element.
	for _, num := range numbers[1:] {
		// Iterates over the remaining numbers in the slice.
		if num < min {
			min = num // Updates the minimum value if a smaller number is found.
		}
	}
	return min // Returns the minimum value.
}

func averageFunc(numbers ...int) int {
	// Calculates the average of the input numbers.
	total := 0 // Initializes the total sum of numbers to 0.
	for _, num := range numbers {
		// Iterates over each number in the input slice.
		total += num // Adds the current number to the total sum.
	}
	// Divides the total sum by the number of elements to find the average.
	return total / len(numbers) // Returns the average value.
}

func maxFunc(numbers ...int) int {
	// Finds the maximum value among the input numbers.
	max := numbers[0] // Initializes the maximum value with the first element.
	for _, num := range numbers[1:] {
		// Iterates over the remaining numbers in the slice.
		if num > max {
			max = num // Updates the maximum value if a larger number is found.
		}
	}
	return max // Returns the maximum value.
}

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func animal(animalType string) (func(float64) float64, string) {
	switch animalType {
	case dog:
		return dogFunc, ""
	case cat:
		return catFunc, ""
	case hamster:
		return hamsterFunc, ""
	case tarantula:
		return tarantulaFunc, ""
	default:
		return nil, "Animal not found"
	}
}

func dogFunc(quantity float64) float64 {
	return quantity * 10
}
func catFunc(quantity float64) float64 {
	return quantity * 5
}
func hamsterFunc(quantity float64) float64 {
	return quantity * 0.25
}
func tarantulaFunc(quantity float64) float64 {
	return quantity * 0.15
}
