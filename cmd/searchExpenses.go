/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	"github.com/spf13/cobra"
)

// searchExpensesCmd represents the searchExpenses command
var searchExpensesCmd = &cobra.Command{
	Use:   "searchExpenses <description>",
	Short: "Search and list expenses by a match in description",
	Args: cobra.ExactArgs(1),
	Aliases: []string{"search"},
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		expenses, err := searchExpenses(args[0])
		if err != nil {
			return err
		}
		printExpensesTable(expenses)

		return nil
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(searchExpensesCmd)
}