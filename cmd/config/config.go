package config

import (
	"github.com/sebcej/githis/cmd/config/set"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure command defaults and filters",
}

func init() {
	ConfigCmd.AddCommand(set.SetCmd)
	ConfigCmd.AddCommand(delCmd)
}
