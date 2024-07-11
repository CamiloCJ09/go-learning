// Package main defines the entry point for the application and includes various function examples.
package main

import (
	"fmt" // Import the fmt package for formatted I/O operations.
)

// main is the entry point of the application.
func main() {
	// Function calls to various examples. Uncomment the desired function call to execute.
	// conditional() // Demonstrates a simple if-else conditional.
	// loop() // Demonstrates a simple for loop.
	// loopWithBreak() // Demonstrates a for loop with a break statement.
	// loopWithContinue() // Demonstrates a for loop with a continue statement.
	// loopWithRange() // Demonstrates a for loop using range (function not implemented in the provided code).
	// wordLenghtAndLetter() // Demonstrates a function to calculate word length and letters (function not implemented in the provided code).
	//employeesOlderThan21() // Calls a function to filter employees older than 21 (function not implemented in the provided code).
	//addEmployee("Benjamin", 20) // Calls a function to add an employee to the map.
	//removeEmployee("Benjamin") // Calls a function to remove an employee from the map.
}

// conditional demonstrates a simple if-else conditional structure.
func conditional() {
	x := 10 // Initialize x with a value of 10.
	if x > 5 {
		fmt.Println("x is greater than 5") // If x is greater than 5, print this message.
	} else {
		fmt.Println("x is less than or equal to 5") // Otherwise, print this message.
	}
}

// loop demonstrates a simple for loop.
func loop() {
	for i := 0; i < 5; i++ { // Loop from 0 to 4.
		fmt.Println(i) // Print the current value of i.
	}
}

// loopWithBreak demonstrates a for loop that exits when a certain condition is met.
func loopWithBreak() {
	for i := 0; i < 5; i++ { // Loop from 0 to 4.
		if i == 3 {
			break // If i is equal to 3, exit the loop.
		}
		fmt.Println(i) // Print the current value of i.
	}
}

// loopWithContinue demonstrates a for loop that skips the rest of the loop body when a certain condition is met.
func loopWithContinue() {
	for i := 0; i < 5; i++ { // Loop from 0 to 4.
		if i == 3 {
			continue // If i is equal to 3, skip the rest of the loop body.
		}
		fmt.Println(i) // Print the current value of i.
	}
}

// Note: The functions loopWithRange, wordLenghtAndLetter, and employeesOlderThan21 are mentioned but not implemented in the provided code.

// loopWithRange iterates over a slice of integers using the range clause.
// It prints the index and value of each element in the slice.
func loopWithRange() {
	nums := []int{2, 3, 4}     // Initialize a slice of integers.
	for i, num := range nums { // Use range to iterate over the slice.
		fmt.Println(i, num) // Print the index and value of each element.
	}
}

// wordLenghtAndLetter reads a word from standard input, prints its length,
// and then prints each letter of the word with its index.
func wordLenghtAndLetter() {
	var word string               // Declare a variable to store the input word.
	fmt.Scanf("%s", &word)        // Read a word from standard input.
	fmt.Println(len(word))        // Print the length of the word.
	for i, letter := range word { // Use range to iterate over the string.
		fmt.Println(i, string(letter)) // Print the index and letter of each character in the word.
	}
}

// employees is a global map storing employee names as keys and their ages as values.
var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// employeesOlderThan21 counts and prints the number of employees older than 21.
func employeesOlderThan21() {
	x := 0                          // Initialize a counter.
	for _, age := range employees { // Iterate over the map using range.
		if age > 21 { // Check if the age is greater than 21.
			x++ // Increment the counter if the condition is true.
		}
	}
	fmt.Println(x) // Print the count of employees older than 21.
}

// addEmployee adds a new employee to the employees map if they do not already exist.
func addEmployee(name string, age int) {
	if _, ok := employees[name]; ok { // Check if the employee already exists.
		fmt.Println("Employee already exists") // Print a message if the employee exists.
		return                                 // Exit the function.
	}
	employees[name] = age                   // Add the new employee to the map.
	fmt.Printf("Employee %s added\n", name) // Print a message confirming the addition.
}

// removeEmployee removes an employee from the employees map if they exist.
func removeEmployee(name string) {
	if _, ok := employees[name]; !ok { // Check if the employee does not exist.
		fmt.Println("Employee doesn't exist") // Print a message if the employee does not exist.
		return                                // Exit the function.
	}
	delete(employees, name) // Remove the employee from the map.
}

// printMonthDayBased prints the month name corresponding to a given number.
func printMonthDayBased(number int) {
	months := map[int]string{ // Initialize a map of month numbers to names.
		1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November", 12: "December",
	}
	result := months[number] // Get the month name from the map using the given number.
	fmt.Println(result)      // Print the month name.
}
