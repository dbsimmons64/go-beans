package main

import (
	"fmt"
	"net/url"
	"strconv"
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

// Return the current year as an int passed as a query parameter.
// If no year parameter is passed then return the current year.
func getYear(query url.Values) int {
	year, err := strconv.Atoi(query.Get("year"))
	if err != nil {
		year = time.Now().Year()
	}

	return year
}

// Return the integer equivalent of the current month string passed as a
// query parameter. If no parameter is passed then set the month to the
// current month
func getMonth(query url.Values) int {

	t, err := time.Parse("Jan", query.Get("month"))
	if err == nil {
		return int(t.Month())
	} else {
		return int(time.Now().Month())
	}
}
