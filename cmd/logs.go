package cmd

import (
	"github.com/spf13/cobra"
)

var (
	day     string
	fromDay string
	toDay   string
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Aggregate and show the commits performed in specified timeframe",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	logsCmd.PersistentFlags().StringVar(&day, "day", "", "Filter by day")

	logsCmd.PersistentFlags().StringVar(&fromDay, "fromDay", "", "Filter start date")
	logsCmd.PersistentFlags().StringVar(&toDay, "toDay", "", "Filter end date")
}
