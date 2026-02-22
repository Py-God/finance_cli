package service

import (
	"encoding/json"
	"github.com/Py-God/finance_cli/models"

	"fmt"
	"log"
	"os"
)


// helper function to read the category json file
func ReadCategoryJson(filename string) ([]models.Category, error) {
	data_bytes, err := os.ReadFile(filename)
	if err != nil {
		return []models.Category{}, err
	}

	// create a list of category struct
	var categories []models.Category
	if len(data_bytes) > 0 {
		err = json.Unmarshal(data_bytes, &categories)
		if err != nil {
			return []models.Category{}, err
		}
	}

	return categories, nil
}

// helper function to write a new category struct (list of structs) object as json
func WriteToCategoryJson(filename string, categories []models.Category) error {
	json_categories, err := json.MarshalIndent(categories, "", "    ")
	if err != nil {
		return fmt.Errorf("Error occurred during marshalling: %w", err)
	}

	// write to the file
	err = os.WriteFile(filename, json_categories, 0644)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	log.Printf("Category successfully Created")

	return nil
}

// helper function to check if the shorthand version of category supplied actually exists
func GetValidCategory(category string) (string, error) {
	filename := "./json/categories.json"
	categories, err := ReadCategoryJson(filename)
	if err != nil {
		return "", err
	}

	for i := range categories {
		if categories[i].Short == category {
			return categories[i].Name, nil
		} else if categories[i].Name == category {
			return categories[i].Name, nil
		}
	}

	return "", fmt.Errorf("the category does not exist")

}

// helper function to read the expense json file
func ReadExpenseJson(filename string) ([]models.Expense, error) {
	data_bytes, err := os.ReadFile(filename)
	if err != nil {
		return []models.Expense{}, err
	}

	// create a list of expenses struct
	var expenses []models.Expense
	if len(data_bytes) > 0 {
		err = json.Unmarshal(data_bytes, &expenses)
		if err != nil {
			return []models.Expense{}, err
		}
	}

	return expenses, nil
}

// helper function to write a new category struct object (list of structs) as json
func WriteToExpenseJson(filename string, expenses []models.Expense, action string) error {
	json_expenses, err := json.MarshalIndent(expenses, "", "    ")
	if err != nil {
		return fmt.Errorf("Error occurred during marshalling: %w", err)
	}

	// write to the file
	err = os.WriteFile(filename, json_expenses, 0644)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	var action_interpreted string
	switch action {
	case "add":
		action_interpreted += "added"
	case "update":
		action_interpreted += "updated"
	case "delete":
		action_interpreted += "deleted"
	}

	log.Printf("Expense successfully %s", action_interpreted)

	return nil
}

// helper function to Create either category or expense json file if they don't exist
func CreateFileIfNotExists(filename string) ([]byte, error) {
	data_bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
		log.Printf("No file named %s is found", filename)
		log.Printf("Creating a new file named %s", filename)
		log.Println("")

		f, err := os.OpenFile(filename, os.O_CREATE, 0644)
		if err != nil {
			return []byte{}, fmt.Errorf("failed to create JSON file: %w", err)
		}

		log.Println("New file created.")
		log.Println("Preparing JSON file...")
		f.Close()
	}

	return data_bytes, err
}