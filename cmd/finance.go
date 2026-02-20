// build packages using
// go build -o bin/finance
// run using ./bin/finance <subcommand>
// cobra-cli add list_expenses

package cmd

import (
	"github.com/google/uuid"
	"encoding/json"
	
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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

func get_valid_category(category string) string {
	var category_interpreted string
	switch category {
	case "t": category_interpreted = "Transportation"
	case "f": category_interpreted = "Feeding"
	case "g": category_interpreted = "Groceries"
	case "m": category_interpreted = "Miscellaneous"
	default:  category_interpreted = category
	}

	return category_interpreted
}

func read_json(filename string) ([]Expense, error) {
	data_bytes, err := os.ReadFile(filename)
	if err != nil {
		return []Expense{}, err
	}

	// create a list of expenses struct
	var expenses []Expense
	if len(data_bytes) > 0{
		err = json.Unmarshal(data_bytes, &expenses)
		if err != nil {
			return []Expense{}, err
		}
	}

	return expenses, nil
}

func add_expense(category string, amount string, description string) (string, error) {
	amount_conv, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", fmt.Errorf("invalid amount: %s", amount)
	}

	category = get_valid_category(category)

	log.Println("Creating Expense...")
	expense := Expense{
		ID:          uuid.New(),
		Category:    category,
		Amount:      amount_conv,
		Description: description,
		Time:        time.Now(),
	}
	log.Println("Expense Created!")

	log.Println("Preparing JSON file...")

	// create file to write to if it doesn't exist, else create it
	filename := "expenses.json"
	data_bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Creating a new file named %s", filename)
		log.Println("")

		f, err := os.OpenFile(filename, os.O_CREATE, 0644)
		if err != nil {
			return "", fmt.Errorf("failed to create JSON file: %w", err)
		}

		log.Println("New file created.")
		log.Println("Preparing JSON file...")
		f.Close()
	}

	
	// create a list of expenses struct
	var expenses []Expense
	// if there is content in json file, unmarshal it, and use that to replace expensese.
	// that's why we used pointers
	if len(data_bytes) > 0{
		err = json.Unmarshal(data_bytes, &expenses)
		if err != nil {
			return "", fmt.Errorf("failed to read JSON file: %w", err)
		}
	}

	// append the new user to expenses
	expenses = append(expenses, expense)

	// Convert users instance to JSON format
	// i guess json can marshal and unmarshal list of structs
    json_expenses, err := json.MarshalIndent(expenses, "", "    ")
    if err != nil {
		return "", fmt.Errorf("Error occurred during marshalling: %w", err)
    }
	fmt.Println("JSON object created successfully")

	// write to the file
	err = os.WriteFile(filename, json_expenses, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return string(data_bytes), nil
}


func list_expenses() (string, error) {
	filename := "expenses.json"
	expenses, err := read_json(filename)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	
	var listed_expenses_string []byte
	if len(expenses) == 0 {
		return "No results found", nil
	} else {
		listed_expenses_string, err = json.MarshalIndent(expenses, "", "    ")
		if err != nil {
			return "", fmt.Errorf("Error occurred during marshalling expenses to display as json: %w", err)
		}
	}

	return string(listed_expenses_string), nil
}


func update_expense(id string, updates map[string]string) error {
	filename := "expenses.json"
	expenses, err := read_json(filename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	id_found := false
	for expense := range(expenses) {
		if expenses[expense].ID.String() == id {
			id_found = true
		}
	}

	if id_found == false {
		return fmt.Errorf("Invalid ID")
	}

	for expense := range(expenses) {
		if expenses[expense].ID.String() == id {
			if updates["category"] != "" {
				expenses[expense].Category = get_valid_category(updates["category"])
			}
			if updates["amount"] != "" {
				amount_int, err := strconv.Atoi(updates["amount"])
				if err != nil {
					return err
				}
				expenses[expense].Amount = float64(amount_int)
			}
			if updates["description"] != "" {
				expenses[expense].Description = updates["description"]
			}
		}
	}

	json_expenses, err := json.MarshalIndent(expenses, "", "    ")
    if err != nil {
		return fmt.Errorf("Error occurred during marshalling: %w", err)
    }
	fmt.Println("JSON object created successfully")

	// write to the file
	err = os.WriteFile(filename, json_expenses, 0644)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	log.Println("Expense successfully updated")

	return nil
}