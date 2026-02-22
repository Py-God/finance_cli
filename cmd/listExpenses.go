/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Py-God/finance_cli/models"
	"github.com/spf13/cobra"
	"github.com/jedib0t/go-pretty/v6/table"
	"encoding/json"
)

var (
	listCategory string
	listDay      string
	listWeek     string
	listMonth    string
	listYear     string
)

// listExpensesCmd represents the listExpenses command
var listExpensesCmd = &cobra.Command{
	Use:     "listExpenses",
	Short:   "List added expenses",
	Long: `Can only use category only, one of the date format only, or both category and one of the date format`,
	Aliases: []string{"list"},
	Args:    cobra.RangeArgs(0, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// use a map to store every flag that was changed
		filters := make(map[string]string)

		if cmd.Flags().Changed("category") {
			filters["category"] = listCategory
		}
		if cmd.Flags().Changed("day") {
			filters["day"] = listDay
		}
		if cmd.Flags().Changed("week") {
			filters["week"] = listWeek
		}
		if cmd.Flags().Changed("month") {
			filters["month"] = listMonth
		}
		if cmd.Flags().Changed("year") {
			filters["year"] = listYear
		}

		expenses, err := listExpenses(filters)
		if err != nil {
			return err
		}
		printExpensesTable(expenses)

		return nil
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(listExpensesCmd)

	listExpensesCmd.Flags().StringVarP(&listCategory, "category", "c", "", "filter by category")
	listExpensesCmd.Flags().StringVarP(&listDay, "day", "d", "", "filter by day (YYYY-MM-DD, DD-MM-YYYY, or DD/MM/YYYY)")
	listExpensesCmd.Flags().StringVarP(&listWeek, "week", "w", "", "filter by week (W or YYYY-W)")
	listExpensesCmd.Flags().StringVarP(&listMonth, "month", "m", "", "filter by month (MM, MonthName, or YYYY-MM)")
	listExpensesCmd.Flags().StringVarP(&listYear, "year", "y", "", "filter by year (YYYY)")
}

// print a nice table specific to expenses struct
func printExpensesTable(expensesString string) {
	var expenses []models.Expense
	json.Unmarshal([]byte(expensesString), &expenses)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Amount", "Description", "Time"})
	
	for _, u := range expenses {
		t.AppendRow(table.Row{u.ID.String(), u.Amount, u.Description, u.Time.String()})
	}
	
	t.Render()
}