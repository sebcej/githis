package cmd

import (
	"fmt"
	"os"

	cfg "github.com/sebcej/githis/cmd/config"
	"github.com/sebcej/githis/cmd/source"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "githis",
	Short: "Aggregate all your project commits",
	Long: `GitHis is a CLI that allows to aggregate and manage all your projects commits in one place.
Start by adding a Projects folder with githis source add [source_folder]`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.githis.yaml)")

	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(source.SourceCmd)
	rootCmd.AddCommand(cfg.ConfigCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".githis" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".githis")
		viper.SafeWriteConfig()
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && cfgFile != "" {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
