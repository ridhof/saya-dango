package main

import (
	"fmt"
	"time"
)


func getMonth(month int32) time.Month {
	switch month {
			case 1:
					return time.January
			case 2:
					return time.February
			case 3:
					return time.March
			case 4:
					return time.April
			case 5: 
					return time.May
			case 6:
					return time.June
			case 7:
					return time.July
			case 8:
					return time.August
			case 9:
					return time.September
			case 10:
					return time.October
			case 11:
					return time.November
			default:
					return time.December
	}
}

func getMonthNumber(month time.Month) int32 {
	switch month {
			case time.January:
					return 1
			case time.February:
					return 2
			case time.March:
					return 3
			case time.April:
					return 4
			case time.May: 
					return 5
			case time.June:
					return 6
			case time.July:
					return 7
			case time.August:
					return 8
			case time.September:
					return 9
			case time.October:
					return 10
			case time.November:
					return 11
			default:
					return 12
	}
}


func libraryFine(dayReturn, monthReturn, yearReturn, dayDue, monthDue, yearDue int32) int32 {
	dueDate := time.Date(int(yearDue), getMonth(monthDue), int(dayDue), 0, 0, 0, 0, time.UTC)
	dateNow := time.Date(int(yearReturn), getMonth(monthReturn), int(dayReturn), 0, 0, 0, 0, time.UTC)

	nextMonth := time.Date(int(yearDue), getMonth(monthDue), 1, 0, 0, 0, 0, time.UTC)
	nextMonth = nextMonth.AddDate(0, 1, 0)
	if dateNow.After(nextMonth) || dateNow.Equal(nextMonth) {
		return 10000
}

	fine := 15 * int32((dateNow.Day() - dueDate.Day()))
	return fine
}

func main() {
	// to do: Running Library Fine Calculation Program 
	fmt.Println("Library Fine Calculation Program")
	fmt.Println(libraryFine(14, 7, 2018, 5, 7, 2018))
	fmt.Println(libraryFine(9, 6, 2015, 6, 6, 2015))
	fmt.Println(libraryFine(9, 7, 2015, 6, 6, 2015))
}