/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addExpenseCmd represents the addExpense command
var addExpenseCmd = &cobra.Command{
	Use:     "addExpense <category> <amount> <description>",
	Aliases: []string{"add"},
	Short:   "Add an expense",
	Args:    cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := add_expense(args[0], args[1], args[2])
		if err != nil {
			return err // Cobra prints "Error: <your message>"
		}
		fmt.Println("Expense added successfully!")
		return nil
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(addExpenseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addExpenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addExpenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
