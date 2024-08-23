package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var returnRecords []Record
	for _, record := range in {
		if predicate(record) {
			returnRecords = append(returnRecords, record)
		}
	}
	return returnRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		for i := p.From; i <= p.To; i++ {
			if record.Day == i {
				return true
			}
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		if record.Category == c {
			return true
		} else {
			return false
		}
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	totalExpense := 0.0
	for day := p.From; day <= p.To; day++ {
		for _, record := range in {
			if record.Day == day {
				totalExpense += record.Amount
			}
		}
	}
	return totalExpense
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	totalExpense := 0.0
	categoryExists := func() bool {
		for _, record := range in {
			if record.Category == c {
				return true
			}
		}
		return false
	}
	if categoryExists() {
		filteredRecords := Filter(Filter(in, ByDaysPeriod(p)), ByCategory(c))
		for _, record := range filteredRecords {
			totalExpense += record.Amount
		}
		return totalExpense, nil
	} else {
		return 0.0, fmt.Errorf("unknown category %s", c)
	}
}
