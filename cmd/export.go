/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export <csv filename>",
	Short: "Export your expenses json to csv",
	Args: cobra.ExactArgs(1),
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return export(args[0])
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
