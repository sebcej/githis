package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/sebcej/githis/aggregator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config = aggregator.Config{}

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Aggregate and show the commits performed in specified timeframe",
	Run: func(cmd *cobra.Command, args []string) {
		sources := []aggregator.Source{}

		viper.UnmarshalKey("sources", &sources)

		config.Filters = aggregator.Filters{}
		logs := aggregator.GetLogs(sources, config, args)

		fmt.Println("Total logs: ", len(logs), "\n")

		if config.Raw {
			json, _ := json.MarshalIndent(logs, "", "    ")

			fmt.Println(string(json))
			return
		}

		makeTable(logs)
	},
}

func makeTable(logs []aggregator.Log) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	tbl := table.New("Hash", "Project", "Author", "Timestamp", "Message")
	tbl.WithHeaderFormatter(headerFmt)

	for _, log := range logs {
		tbl.AddRow(log.Hash, log.Project, log.Author.Name, log.Date, log.Message)
	}

	tbl.Print()
}

func init() {
	logsCmd.PersistentFlags().IntVar(&config.Offset, "offset", 0, "Set positive or negative offset based on today")
	logsCmd.PersistentFlags().BoolVar(&config.FullMessage, "fullMessage", false, "Show full commit messages")
	logsCmd.PersistentFlags().BoolVar(&config.Raw, "raw", false, "Show RAW json git output")
	logsCmd.PersistentFlags().StringVar(&config.FromDay, "fromDay", "", "Start date for commit filter")
	logsCmd.PersistentFlags().StringVar(&config.ToDay, "toDay", "", "End date for commit filter")
}
