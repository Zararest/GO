package main

import (
	"encoding/csv"
	"example/pkg/fib"
	"fmt"
	"os"
	"strconv"
)

func foo() {
	panicVar = false
}

var panicVar = true

type Student struct {
	Name    string
	Grades  []int
	Average float64
}

func calculateAverage(grades []int) float64 {
	total := 0
	for _, grade := range grades {
		total += grade
	}
	return float64(total) / float64(len(grades))
}

func main() {

	if panicVar {
		panic("AAAA")
	} else {
		fmt.Println("Everything is fine")
	}

	fib.FibMain()

	// Open the input CSV file
	inputFile, err := os.Open("grades.csv")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(inputFile)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	// Parse CSV records into Student structs
	var students []Student
	for _, record := range records {
		grades := make([]int, len(record)-1)
		for i, gradeStr := range record[1:] {
			grade, err := strconv.Atoi(gradeStr)
			if err != nil {
				fmt.Println("Error parsing grade:", err)
				return
			}
			grades[i] = grade
		}
		average := calculateAverage(grades)
		student := Student{Name: record[0], Grades: grades, Average: average}
		students = append(students, student)
	}

	// Open the output CSV file
	outputFile, err := os.Create("averages.csv")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write headers to the output CSV file
	if err := writer.Write([]string{"Name", "Average Grade"}); err != nil {
		fmt.Println("Error writing headers to output file:", err)
		return
	}

	// Write student averages to the output CSV file
	for _, student := range students {
		record := []string{student.Name, fmt.Sprintf("%.2f", student.Average)}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to output file:", err)
			return
		}
	}

	fmt.Println("Averages written to averages.csv successfully.")
}
