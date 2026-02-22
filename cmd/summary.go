/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	summaryDay      string
	summaryWeek     string
	summaryMonth    string
	summaryYear     string
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary [flag] <day> | <week> | <month> | <year>",
	Short: "total amount in expenses spent in a period",
	RunE: func(cmd *cobra.Command, args []string) error {
		filters := make(map[string]string)

		if cmd.Flags().Changed("day") {
			filters["day"] = summaryDay
		}
		if cmd.Flags().Changed("week") {
			filters["week"] = summaryWeek
		}
		if cmd.Flags().Changed("month") {
			filters["month"] = summaryMonth
		}
		if cmd.Flags().Changed("year") {
			filters["year"] = summaryYear
		}

		// Apply all updates in one go
		expenseSummary, err := summary(filters)
		if err != nil {
			return err
		}
		fmt.Println(expenseSummary)

		return nil
	},

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().StringVarP(&summaryDay, "day", "d", "", "summarize by day (YYYY-MM-DD, DD-MM-YYYY, or DD/MM/YYYY)")
	summaryCmd.Flags().StringVarP(&summaryWeek, "week", "w", "", "summarize by week (W or YYYY-W)")
	summaryCmd.Flags().StringVarP(&summaryMonth, "month", "m", "", "summarize by month (MM, MonthName, or YYYY-MM)")
	summaryCmd.Flags().StringVarP(&summaryYear, "year", "y", "", "summarize by year (YYYY)")
}
