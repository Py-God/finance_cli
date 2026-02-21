/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteExpenseCmd represents the deleteExpense command
var deleteExpenseCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete an expense from database",
	Aliases: []string{"del"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := deleteExpense(args[0])

		if err == nil {
			fmt.Println("Expense Successfully Deleted")
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(deleteExpenseCmd)
}
