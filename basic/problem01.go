package main

// Problem: Student Grade Calculator

// Create a Go program that displays a menu with three options:

// Calculate Grade
// Check Pass/Fail
// Exit

// Requirements:

// Ask the user to select an option.
// For Calculate Grade, take marks as input and display the grade.
// For Check Pass/Fail, take marks and display whether the student passed or failed.
// A student passes if the marks are 40 or above.
// Keep showing the menu until the user selects Exit.
// Handle invalid menu choices.

// Grade system:
// 80–100 = A+
// 70–79  = A
// 60–69  = B
// 50–59  = C
// 40–49  = D
// 0–39   = F

// Example Output

// 1. Calculate Grade
// 2. Check Pass/Fail
// 3. Exit

// Enter choice: 1
// Enter marks: 85

// Grade: A+

// Enter choice: 2
// Enter marks: 35

// Result: Fail

// Enter choice: 3

// Goodbye!

import "fmt"

func getGrade(marks int) string {
	switch {
	case marks >= 80 && marks <= 100:
		return "A+"
	case marks >= 70 && marks <= 79:
		return "A"
	case marks >= 60 && marks <= 69:
		return "B"
	case marks >= 50 && marks <= 59:
		return "C"
	case marks >= 40 && marks <= 49:
		return "D"
	case marks >= 0 && marks <= 39:
		return "F"
	default:
		return "Invalid marks"
	}
}

func isPassing(marks int) bool {
	return marks >= 40
}

func main() {
	isRunning := true

	for isRunning {
		var number int
		println("1. Calculate Grade")
		println("2. Check Pass/Fail")
		println("3. Exit")
		fmt.Print("\nEnter choice: ")
		fmt.Scanln(&number)

		switch number {
		case 1:
			var marks int
			fmt.Print("Enter marks: ")
			fmt.Scanln(&marks)
			grade := getGrade(marks)
			fmt.Printf("Grade: %s\n", grade)
		case 2:
			var marks int
			fmt.Print("Enter marks: ")
			fmt.Scanln(&marks)
			if isPassing(marks) {
				fmt.Println("Result: Pass")
			} else {
				fmt.Println("Result: Fail")
			}
		case 3:
			isRunning = false
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}

}
