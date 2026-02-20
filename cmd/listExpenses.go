/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listExpensesCmd represents the listExpenses command
var listExpensesCmd = &cobra.Command{
	Use:   "listExpenses",
	Short: "List added expenses",
	Aliases: []string{"list"},
	RunE: func(cmd *cobra.Command, args []string) error {
		expenses, err := listExpenses()
		if err != nil {
			return err
		}
		fmt.Println(expenses)

		return nil
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(listExpensesCmd)
}
