package main

import "fmt"

// Create a Student structure
type Student struct {
    Name  string
    Grade float64
}

// Function to calculate a letter grade
func getLetterGrade(grade float64) string {
    if grade >= 90 {
        return "A"
    } else if grade >= 80 {
        return "B"
    } else if grade >= 70 {
        return "C"
    } else if grade >= 60 {
        return "D"
    }

    return "F"
}

func main() {
    // Create a slice of students
    students := []Student{
        {"Alex", 95},
        {"Jordan", 82},
        {"Taylor", 73},
        {"Chris", 58},
    }

    // Loop through the students
    for _, student := range students {
        letter := getLetterGrade(student.Grade)

        fmt.Println("Student:", student.Name)
        fmt.Println("Grade:", student.Grade)
        fmt.Println("Letter:", letter)
        fmt.Println("----------------")
    }
}