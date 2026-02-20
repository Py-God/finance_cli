/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
    updateCategory    string
    updateAmount      string
    updateDescription string
)

// updateExpenseCmd represents the updateExpense command
var updateExpenseCmd = &cobra.Command{
	Use:   "finance_cli update <expense_uuid> -c <category> -a <amount> -d <description>",
	Short: "A brief description of your command",
	Aliases: []string{"update"},
	Args: cobra.MinimumNArgs(1),
	Long: `if using all three flags, ensure to use in that format, else, ensure the flags follow that sequence.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		updates := make(map[string]string)
    
		if cmd.Flags().Changed("category") {
			updates["category"] = updateCategory
		}
		if cmd.Flags().Changed("amount") {
			updates["amount"] = updateAmount
		}
		if cmd.Flags().Changed("description") {
			updates["description"] = updateDescription
		}
		
		// Validate 1-3 fields provided
		if len(updates) == 0 || len(updates) > 3 {
			return fmt.Errorf("please provide between 1 and 3 fields to update")
		}
		
		// Apply all updates in one go
		return update_expense(args[0], updates)
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(updateExpenseCmd)

	updateExpenseCmd.Flags().StringVarP(&updateCategory, "category", "c", "", "update category")  
	updateExpenseCmd.Flags().StringVarP(&updateAmount, "amount", "a", "", "update amount")  
	updateExpenseCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "update description")
}
