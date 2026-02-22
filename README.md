# Finance CLI

A command-line expense tracker to manage your finances efficiently. Track expenses, categorize spending, and generate summaries with ease.

https://roadmap.sh/projects/expense-tracker

## Features

- **Add Categories**: Create custom expense categories
- **Add Expenses**: Log expenses with category, amount, and description
- **List Expenses**: View all expenses with filtering options (by category, day, week, month, year)
- **Search Expenses**: Search expenses by description
- **Update Expenses**: Modify expense details (category, amount, description)
- **Delete Expenses**: Remove expenses from your records
- **Summary**: Get spending summaries for specific time periods
- **Export**: Export expenses to CSV format

## Project Structure
```
finance_cli/
├── main.go              # Entry point
├── cmd/                 # Command implementations
│   ├── root.go         # Root command
│   ├── main.go         # Helper functions
│   ├── addCategory.go
│   ├── addExpense.go
│   ├── listExpenses.go
│   ├── searchExpenses.go
│   ├── updateExpense.go
│   ├── deleteExpense.go
│   ├── categories.go
│   ├── summary.go
│   └── export.go
├── models/             # Data structures
├── service/            # Business logic
├── utils/              # Utility functions
├── json/               # JSON data files
├── exports/            # CSV exports
├── bin/                # Compiled binaries
├── go.mod              # Module definition
└── LICENSE             # License file
```
## Installation

### Prerequisites
- Go 1.20 or higher

### Build
```sh
go build -o bin/finance
```

## Usage

```sh
./bin/finance [command] [options]
```

### Commands

- `add-category <name>` - Create a new expense category
- `add-expense <category> <amount> <description>` - Log a new expense
- `list` - Display all expenses with optional filters
- `search <keyword>` - Find expenses by description
- `update <id> [category] [amount] [description]` - Modify an expense
- `delete <id>` - Remove an expense
- `summary [period]` - Show spending summary (day/week/month/year)
- `export <filename>` - Save expenses to CSV file

## Examples

```sh
# Add a category
./bin/finance add-category "Groceries"

# Log an expense
./bin/finance add-expense "Groceries" 50.25 "Weekly shopping"

# View all expenses
./bin/finance list

# Get monthly summary
./bin/finance summary month

# Export to CSV
./bin/finance export expenses.csv
```

## License

MIT License
