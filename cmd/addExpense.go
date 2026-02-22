/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addExpenseCmd represents the addExpense command
var addExpenseCmd = &cobra.Command{
	Use:     "addExpense <category> <amount> <description>",
	Aliases: []string{"add"},
	Short:   "Add an expense. Ensure you have added a category first.",
	Args:    cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := addExpense(args[0], args[1], args[2])
		if err != nil {
			return err
		}
		
		return nil
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(addExpenseCmd)
}
