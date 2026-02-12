// build packages using
// go build -o bin/finance
// run using ./bin/finance <subcommand>

package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/google/uuid"
	"encoding/json"
)


type Expense struct {
	ID          uuid.UUID `json:"id"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}


func display_categories() string {
	message := `Available Fields (not case sensitive)
	Transportation - t
	Feeding - f
	Groceries - g
	Miscellaneous - m`

	return message
}

func add_expense(category string, amount string, description string) (string, error) {
	amount_conv, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", fmt.Errorf("invalid amount: %s", amount)
	}

	var category_interpreted string
	switch category {
	case "t": category_interpreted = "Transportation"
	case "f": category_interpreted = "Feeding"
	case "g": category_interpreted = "Groceries"
	case "m": category_interpreted = "Miscellaneous"
	default:  category_interpreted = "Other"
	}

	expense := Expense{
		ID:          uuid.New(),
		Category:    category_interpreted,
		Amount:      amount_conv,
		Description: description,
		Time:        time.Now(),
	}

	filename := "expenses.json"
	
	// Read existing data
	var data []Expense
	file, err := os.ReadFile(filename)
	if err == nil && len(file) > 0 {
		json.Unmarshal(file, &data)
	}

	data = append(data, expense)

	// Use MarshalIndent directly on the data object
	dataBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %w", err)
	}

	err = os.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return string(dataBytes), nil
}
