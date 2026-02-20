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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteExpenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteExpenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
