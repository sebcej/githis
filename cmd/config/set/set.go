package set

import (
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config params",
}

func init() {
	SetCmd.AddCommand(authorCmd)
}
