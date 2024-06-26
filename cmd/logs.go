package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/sebcej/githis/aggregator"
	"github.com/sebcej/githis/out"
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

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Prefix = "\n"
		s.Suffix = "\n"
		s.Color("green")

		// Set author default if available
		if len(config.Filters.Authors) == 0 {
			authorFilter := viper.GetString("author")

			if authorFilter != "" {
				fmt.Println("By author: ", authorFilter)
				config.Filters.Authors = []string{authorFilter}
			}
		}

		s.Start()

		logs, logsLen := aggregator.GetLogs(sources, config, args)

		s.Stop()

		fmt.Printf("Total logs: %d\n\n", logsLen)

		if config.Raw {
			json, _ := json.MarshalIndent(logs, "", "    ")

			fmt.Println(string(json))
			return
		}

		out.MakeStatic(logs)
	},
}

func init() {
	logsCmd.PersistentFlags().BoolVar(&config.FullMessage, "fullMessage", false, "Show full commit messages")
	logsCmd.PersistentFlags().BoolVar(&config.Raw, "raw", false, "Show RAW json git output")
	logsCmd.PersistentFlags().BoolVarP(&config.Pull, "pull", "p", false, "Pull the repo before")
	logsCmd.PersistentFlags().BoolVarP(&config.Reverse, "reverse", "r", false, "Reverse the order of logs (from older to newer)")

	logsCmd.PersistentFlags().StringSliceVarP(&config.Filters.Authors, "author", "a", []string{}, "Filter by commit author")
	logsCmd.PersistentFlags().IntVarP(&config.Filters.Limit, "limit", "l", 100, "Limit number of show commits")
	logsCmd.PersistentFlags().IntVarP(&config.Filters.Offset, "offset", "o", 0, "Set positive or negative offset based on today")
	logsCmd.PersistentFlags().StringVarP(&config.Filters.Day, "day", "d", "", "Filter by single day")
	logsCmd.PersistentFlags().StringVar(&config.Filters.From, "from", "", "Start date for commit filter")
	logsCmd.PersistentFlags().StringVar(&config.Filters.To, "to", "", "End date for commit filter")
}
