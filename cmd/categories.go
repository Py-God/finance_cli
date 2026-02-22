/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/Py-God/finance_cli/models"
	"github.com/jedib0t/go-pretty/v6/table"
	"encoding/json"
	"github.com/spf13/cobra"
)

// categoriesCmd represents the categories command
var ListcategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Display available categories",
	RunE: func(cmd *cobra.Command, args []string) error {
		categories, err := ListCategories()
		if err != nil {
			return err
		}
		printCategoriesTable(categories)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ListcategoriesCmd)
}

// print a nice table specific for the categories struct
func printCategoriesTable(categoriesString string) {
	var categories []models.Category
	json.Unmarshal([]byte(categoriesString), &categories)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Short", "Name"})
	
	for _, u := range categories {
		t.AppendRow(table.Row{u.Short, u.Name})
	}
	
	t.Render()
}