/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	category_given bool
	date_given bool
)

// listExpensesCmd represents the listExpenses command
var listExpensesCmd = &cobra.Command{
	Use:   "listExpenses",
	Short: "List added expenses",
	Long: `flag --category or -c <category> to filter using category
flag --date or  -d to filter using date
no flags to list all available expenses`,
	Aliases: []string{"list"},
	Args: cobra.RangeArgs(0, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var expenses string
		if category_given && len(args) > 0{
			expenses, err = list_expenses(args[0], "")
		} else if date_given && len(args) > 0 {
			expenses, err = list_expenses("", args[0])
		} else if category_given && date_given && len(args) > 0 {
			expenses, err = list_expenses(args[0], args[1])
		} else {
			if len(args) == 0 {
				fmt.Println("No arguments were given.")
			}
		}

		if err != nil {
			return err // Cobra prints "Error: <your message>"
		}
		fmt.Println(expenses)

		return nil
	},

	SilenceUsage: true,
}

func init() {
	listExpensesCmd.Flags().BoolVarP(&category_given, "category", "c", false, "category is supplied to list_expenses")
	listExpensesCmd.Flags().BoolVarP(&date_given, "date", "d", false, "date is supplied to list_expenses")  

	rootCmd.AddCommand(listExpensesCmd)
}
