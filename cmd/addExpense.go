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
		err := addExpense(args[0], args[1], args[2])
		if err != nil {
			return err
		}
		fmt.Println("Expense added successfully!")
		return nil
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(addExpenseCmd)
}
