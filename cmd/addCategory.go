/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCategoryCmd represents the addCategory command
var addCategoryCmd = &cobra.Command{
	Use:   "addCategory <short name of category> <full name of category>",
	Short: "Add a category that can be used when adding an expense",
	Long: `E.g.
f for Feeding
g for Groceries`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := addCategory(args[0], args[1])

		if err == nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCategoryCmd)
}
