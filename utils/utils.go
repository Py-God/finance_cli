package utils

import (
	"github.com/Py-God/finance_cli/models"

	"fmt"
	"time"
)

// helper function delete an expense, and reorder the index.
func DeleteAtIndex(slice []models.Expense, index int) []models.Expense {

	// Append function used to append elements to a slice
	// first parameter as the slice to which the elements
	// are to be added/appended second parameter is the
	// element(s) to be appended into the slice
	// return value as a slice
	return append(slice[:index], slice[index+1:]...)
}

// helper function to check if supplied day is of valid format
func ParseDay(s string) (time.Time, error) {
	layouts := []string{"2006-01-02", "02-01-2006", "02/01/2006"}
	var t time.Time
	var err error
	for _, l := range layouts {
		t, err = time.Parse(l, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not parse day: %s", s)
}