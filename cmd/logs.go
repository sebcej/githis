package cmd

import (
	"fmt"

	"github.com/sebcej/githis/aggregator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		sources := []aggregator.Source{}

		viper.UnmarshalKey("sources", &sources)

		filters := aggregator.Filters{}
		logs := aggregator.GetLogs(sources, filters, args)

		fmt.Println("Logs", logs)
	},
}

func init() {
	logsCmd.PersistentFlags().StringVar(&day, "day", "", "Filter by day")

	logsCmd.PersistentFlags().StringVar(&fromDay, "fromDay", "", "Filter start date")
	logsCmd.PersistentFlags().StringVar(&toDay, "toDay", "", "Filter end date")
}
