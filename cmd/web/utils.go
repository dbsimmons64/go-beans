package main

import (
	"fmt"
	"time"
)

type Month struct {
	Name  string
	Value int
}

func ListMonths() []Month {
	months := []Month{}
	for month := time.January; month <= time.December; month++ {
		months = append(months, Month{Name: month.String(), Value: int(month)})
	}

	return months
}

// Returns a list of years to allow selections of transactions from previous years.
func ListYears() []int {
	years := []int{}

	today := time.Now()
	currentYear := today.Year()

	for year := (currentYear - 10); year <= currentYear; year++ {
		years = append(years, year)
	}

	fmt.Println(years)

	return years
}
