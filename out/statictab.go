package out

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
	"github.com/rodaine/table"
	"github.com/sebcej/githis/aggregator"
)

var matchCommitUrl = regexp.MustCompile(`^\033.*\033\\$`)

func MakeStatic(logs []aggregator.Log) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	tbl := table.New("Hash", "Project", "Author", "Timestamp", "Message")
	tbl.WithHeaderFormatter(headerFmt).WithWidthFunc(func(s string) int {
		str := matchCommitUrl.ReplaceAllString(s, "*******")

		return runewidth.StringWidth(str)
	})

	for _, log := range logs {
		hashCol := log.Hash

		if log.CommitUrl != "" {
			hashCol = fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", log.CommitUrl, hashCol)
		}

		tbl.AddRow(hashCol, log.Project, log.Author.Name, log.Date, log.Message)
	}

	tbl.Print()
}
