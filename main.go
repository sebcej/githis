package main

import (
	"os/exec"

	"github.com/sebcej/githis/cmd"
	"github.com/spf13/cobra"
)

func main() {
	_, err := exec.LookPath("git")
	cobra.CheckErr(err)

	cmd.Execute()
}
