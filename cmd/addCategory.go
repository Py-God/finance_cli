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
	Short: "A brief description of your command",
	Long: ``,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCategoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
