package out

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/sebcej/githis/aggregator"
)

func MakeStatic(logs []aggregator.Log) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	tbl := table.New("Hash", "Project", "Author", "Timestamp", "Message")
	tbl.WithHeaderFormatter(headerFmt)

	for _, log := range logs {
		tbl.AddRow(log.Hash, log.Project, log.Author.Name, log.Date, log.Message)
	}

	tbl.Print()
}
