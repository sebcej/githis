package source

import (
	"github.com/spf13/cobra"
)

type Source struct {
	Name string
	Path string
}

var SourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage sources from which the git stats will be extracted",
}

func init() {
	SourceCmd.AddCommand(listCmd)
	SourceCmd.AddCommand(addCmd)
	SourceCmd.AddCommand(removeCmd)
}
