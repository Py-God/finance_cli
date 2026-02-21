// build packages using
// go build -o bin/finance
// run using ./bin/finance <subcommand>
// cobra-cli add list_expenses

package cmd

import (
	"encoding/json"
	"encoding/csv"

	"github.com/google/uuid"

	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	ID          uuid.UUID `json:"id"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

type Category struct {
	Short string `json:"short"`
	Name  string `json:"name"`
}

// helper function to List all the added categories
func ListCategories() (string, error) {
	// Read the category json
	filename := "categories.json"
	categories, err := readCategoryJson(filename)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	// return the category json as a string
	var listed_categories_string []byte
	if len(categories) == 0 {
		return "No results found", nil
	} else {
		listed_categories_string, err = json.MarshalIndent(categories, "", "    ")
		if err != nil {
			return "", fmt.Errorf("Error occurred during marshalling expenses to display as json: %w", err)
		}
	}

	return string(listed_categories_string), nil
}

// helper function to read the category json file
func readCategoryJson(filename string) ([]Category, error) {
	data_bytes, err := os.ReadFile(filename)
	if err != nil {
		return []Category{}, err
	}

	// create a list of category struct
	var categories []Category
	if len(data_bytes) > 0 {
		err = json.Unmarshal(data_bytes, &categories)
		if err != nil {
			return []Category{}, err
		}
	}

	return categories, nil
}

// helper function to check if the shorthand version of category supplied actually exists
func getValidCategory(category string) (string, error) {
	filename := "categories.json"
	categories, err := readCategoryJson(filename)
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
func readExpenseJson(filename string) ([]Expense, error) {
	data_bytes, err := os.ReadFile(filename)
	if err != nil {
		return []Expense{}, err
	}

	// create a list of expenses struct
	var expenses []Expense
	if len(data_bytes) > 0 {
		err = json.Unmarshal(data_bytes, &expenses)
		if err != nil {
			return []Expense{}, err
		}
	}

	return expenses, nil
}

// helper function to write a new category struct (list of structs) object as json
func writeToCategoryJson(filename string, categories []Category) error {
	json_categories, err := json.MarshalIndent(categories, "", "    ")
	if err != nil {
		return fmt.Errorf("Error occurred during marshalling: %w", err)
	}
	fmt.Println("JSON object created successfully")

	// write to the file
	err = os.WriteFile(filename, json_categories, 0644)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	log.Printf("Category successfully Created")

	return nil
}

// helper function to write a new category struct object (list of structs) as json
func writeToExpenseJson(filename string, expenses []Expense, action string) error {
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

// helper function delete an expense, and reorder the index.
func deleteAtIndex(slice []Expense, index int) []Expense {

	// Append function used to append elements to a slice
	// first parameter as the slice to which the elements
	// are to be added/appended second parameter is the
	// element(s) to be appended into the slice
	// return value as a slice
	return append(slice[:index], slice[index+1:]...)
}

// helper function to Create either category or expense json file if they don't exist
func createFileIfNotExists(filename string) ([]byte, error) {
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

// helper function to check if supplied day is of valid format
func parseDay(s string) (time.Time, error) {
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

// add a new category
func addCategory(short string, name string) error {
	log.Println("Creating Category...")
	category := Category{
		Short: short,
		Name:  name,
	}
	log.Println("Category Created!")

	log.Println("Preparing JSON file...")

	// create file to write to if it doesn't exist, else create it
	filename := "categories.json"
	data_bytes, err := createFileIfNotExists(filename)

	// create a list of expenses struct
	var categories []Category
	// if there is content in json file, unmarshal it, and use that to replace expensese.
	// that's why we used pointers
	if len(data_bytes) > 0 {
		err = json.Unmarshal(data_bytes, &categories)
		if err != nil {
			return fmt.Errorf("failed to read JSON file: %w", err)
		}
	}

	// append the new user to expenses
	categories = append(categories, category)

	return writeToCategoryJson(filename, categories)
}

// add a new expense
func addExpense(category string, amount string, description string) error {
	amount_conv, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("invalid amount: %s", amount)
	}

	filename := "categories.json"
	data_bytes, err := createFileIfNotExists(filename)

	category, err = getValidCategory(category)
	if err != nil {
		return err
	}

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
	filename = "expenses.json"
	data_bytes, err = createFileIfNotExists(filename)

	// create a list of expenses struct
	var expenses []Expense
	// if there is content in json file, unmarshal it, and use that to replace expensese.
	// that's why we used pointers
	if len(data_bytes) > 0 {
		err = json.Unmarshal(data_bytes, &expenses)
		if err != nil {
			return fmt.Errorf("failed to read JSON file: %w", err)
		}
	}

	// append the new user to expenses
	expenses = append(expenses, expense)

	return writeToExpenseJson(filename, expenses, "add")
}

// list all added expenses
func listExpenses(filters map[string]string) (string, error) {
	filename := "expenses.json"
	expenses, err := readExpenseJson(filename)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	var expensesFiltered []Expense
	if len(expenses) == 0 {
		return "No Expenses found", nil
	}

	// Got this from copilot
	// Normalize category filter: accept short or full name
	if filters["category"] != "" {
		if catName, err := getValidCategory(filters["category"]); err == nil {
			filters["category"] = catName
		}
	}

	for i := range expenses {
		e := expenses[i]
		match := true

		if filters["category"] != "" {
			if e.Category != filters["category"] {
				match = false
			}
		}

		if match && filters["day"] != "" {
			d, err := parseDay(filters["day"])
			if err != nil {
				return "", err
			}
			if !(e.Time.Year() == d.Year() && e.Time.Month() == d.Month() && e.Time.Day() == d.Day()) {
				match = false
			}
		}

		if match && filters["week"] != "" {
			// Accept either "WEEK" or "YEAR-WEEK"
			fw := filters["week"]
			var wantYear, wantWeek int
			if strings.Contains(fw, "-") {
				parts := strings.Split(fw, "-")
				if len(parts) == 2 {
					wantYear, _ = strconv.Atoi(parts[0])
					wantWeek, _ = strconv.Atoi(parts[1])
				}
			} else {
				wantWeek, _ = strconv.Atoi(fw)
			}
			y, w := e.Time.ISOWeek()
			if wantWeek != 0 && w != wantWeek {
				match = false
			}
			if wantYear != 0 && y != wantYear {
				match = false
			}
		}

		if match && filters["month"] != "" {
			// Accept "MM", "MonthName", or "YYYY-MM"
			fm := filters["month"]
			var wantYear, wantMonth int
			if strings.Contains(fm, "-") {
				parts := strings.Split(fm, "-")
				if len(parts) == 2 {
					wantYear, _ = strconv.Atoi(parts[0])
					wantMonth, _ = strconv.Atoi(parts[1])
				}
			} else {
				// try numeric month
				m, err := strconv.Atoi(fm)
				if err == nil {
					wantMonth = m
				} else {
					// try parse month name
					if t, err := time.Parse("January", fm); err == nil {
						wantMonth = int(t.Month())
					} else if t, err := time.Parse("Jan", fm); err == nil {
						wantMonth = int(t.Month())
					}
				}
			}
			if wantMonth != 0 && int(e.Time.Month()) != wantMonth {
				match = false
			}
			if wantYear != 0 && e.Time.Year() != wantYear {
				match = false
			}
		}

		if match && filters["year"] != "" {
			y, err := strconv.Atoi(filters["year"])
			if err != nil {
				return "", err
			}
			if e.Time.Year() != y {
				match = false
			}
		}

		if match {
			expensesFiltered = append(expensesFiltered, e)
		}
	}

	expensesFilteredBytes, err := json.MarshalIndent(expensesFiltered, "", "    ")
	if err != nil {
		return "", fmt.Errorf("Error occurred during marshalling expenses to display as json: %w", err)
	}

	return string(expensesFilteredBytes), nil
}

// update an expense (either category, amount or description)
func updateExpense(id string, updates map[string]string) error {
	filename := "expenses.json"
	expenses, err := readExpenseJson(filename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	id_found := false
	for expense := range expenses {
		if expenses[expense].ID.String() == id {
			id_found = true
		}
	}

	if id_found == false {
		return fmt.Errorf("Invalid ID")
	}

	for expense := range expenses {
		if expenses[expense].ID.String() == id {
			if updates["category"] != "" {
				expenses[expense].Category, err = getValidCategory(updates["category"])
				if err != nil {
					return err
				}
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

	return writeToExpenseJson(filename, expenses, "update")
}

// delete an expense
func deleteExpense(id string) error {
	filename := "expenses.json"
	expenses, err := readExpenseJson(filename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	id_found := false
	var i int
	for i = range expenses {
		if expenses[i].ID.String() == id {
			id_found = true
			break
		}
	}

	if id_found == false {
		return fmt.Errorf("Invalid ID")
	}

	expenses = deleteAtIndex(expenses, i)

	return writeToExpenseJson(filename, expenses, "delete")
}

// search from added expenses
func searchExpenses(description string) (string, error) {
	filename := "expenses.json"
	expenses, err := readExpenseJson(filename)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	var expensesFiltered []Expense
	if len(expenses) == 0 {
		return "No Expenses found", nil
	}

	for i := range expenses {
		if strings.Contains(strings.ToLower(expenses[i].Description), strings.ToLower(description)) {
			expensesFiltered = append(expensesFiltered, expenses[i])
		}
	}

	expensesFilteredBytes, err := json.MarshalIndent(expensesFiltered, "", "    ")
	if err != nil {
		return "", fmt.Errorf("Error occurred during marshalling expenses to display as json: %w", err)
	}

	return string(expensesFilteredBytes), nil
}

// get an amount summary of the expenses for a period
func summary(filters map[string]string) (string, error) {
	filename := "expenses.json"
	expenses, err := readExpenseJson(filename)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if len(expenses) == 0 {
		return "No Expenses found", nil
	}

	var totalExpenses float64 = 0
	for i := range expenses {
		e := expenses[i]
		match := true

		if match && filters["day"] != "" {
			d, err := parseDay(filters["day"])
			if err != nil {
				return "", err
			}
			if !(e.Time.Year() == d.Year() && e.Time.Month() == d.Month() && e.Time.Day() == d.Day()) {
				match = false
			}
		}

		if match && filters["week"] != "" {
			// Accept either "WEEK" or "YEAR-WEEK"
			fw := filters["week"]
			var wantYear, wantWeek int
			if strings.Contains(fw, "-") {
				parts := strings.Split(fw, "-")
				if len(parts) == 2 {
					wantYear, _ = strconv.Atoi(parts[0])
					wantWeek, _ = strconv.Atoi(parts[1])
				}
			} else {
				wantWeek, _ = strconv.Atoi(fw)
			}
			y, w := e.Time.ISOWeek()
			if wantWeek != 0 && w != wantWeek {
				match = false
			}
			if wantYear != 0 && y != wantYear {
				match = false
			}
		}

		if match && filters["month"] != "" {
			// Accept "MM", "MonthName", or "YYYY-MM"
			fm := filters["month"]
			var wantYear, wantMonth int
			if strings.Contains(fm, "-") {
				parts := strings.Split(fm, "-")
				if len(parts) == 2 {
					wantYear, _ = strconv.Atoi(parts[0])
					wantMonth, _ = strconv.Atoi(parts[1])
				}
			} else {
				// try numeric month
				m, err := strconv.Atoi(fm)
				if err == nil {
					wantMonth = m
				} else {
					// try parse month name
					if t, err := time.Parse("January", fm); err == nil {
						wantMonth = int(t.Month())
					} else if t, err := time.Parse("Jan", fm); err == nil {
						wantMonth = int(t.Month())
					}
				}
			}
			if wantMonth != 0 && int(e.Time.Month()) != wantMonth {
				match = false
			}
			if wantYear != 0 && e.Time.Year() != wantYear {
				match = false
			}
		}

		if match && filters["year"] != "" {
			y, err := strconv.Atoi(filters["year"])
			if err != nil {
				return "", err
			}
			if e.Time.Year() != y {
				match = false
			}
		}

		if match {
			totalExpenses += e.Amount
		}
	}

	if totalExpenses == 0 {
		return "No expenses for that period", nil
	}
	return fmt.Sprintf("Total expenses for that period: %.2f", totalExpenses), nil
}

// export expenses.json to csv
func export(csvFilename string) error {
	if !strings.HasSuffix(csvFilename, ".csv") {
		return fmt.Errorf("File must be of type csv")
	}
	
	expenses, err := readExpenseJson("expenses.json")
	if err != nil {
		return err
	}

    // 3. Create a new file to store CSV data
    outputFile, err := os.Create(csvFilename)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    // 4. Write the header of the CSV file and the successive rows by iterating through the JSON struct array
    writer := csv.NewWriter(outputFile)
    defer writer.Flush()

    header := []string{"ID", "Category", "Amount", "Description", "Time"}
    if err := writer.Write(header); err != nil {
        return err
    }

    for _, r := range expenses {
        var csvRow []string
        csvRow = append(csvRow, r.ID.String(), r.Category, fmt.Sprintf("%.2f", r.Amount), r.Time.String())
        if err := writer.Write(csvRow); err != nil {
            return err
        }
    }

    return nil
}