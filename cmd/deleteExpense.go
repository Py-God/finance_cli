/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deleteExpenseCmd represents the deleteExpense command
var deleteExpenseCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete an expense from database",
	Aliases: []string{"del"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteExpense(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteExpenseCmd)
}
