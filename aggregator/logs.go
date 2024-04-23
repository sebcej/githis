package aggregator

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func getLogsFromGit(project, dir string, config Config, extraArgs []string) (logs []Log) {
	builtArgs := []string{"log", "--all", "--date=format:%Y-%m-%d %H:%M:%S", commitFormat}
	builtArgs = append(builtArgs, extraArgs...)

	if config.Pull {
		pullCmd := exec.Command("git", "pull", "--quiet")
		pullCmd.Dir = dir

		_, err := pullCmd.CombinedOutput()
		if err != nil {
			return
		}
	}

	cmd := exec.Command("git", builtArgs...)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	output := string(out)
	output = trailingComma.ReplaceAllString(output, "")
	output = strings.ReplaceAll(output, `\`, `\\`)
	output = strings.ReplaceAll(output, `"`, `\"`)
	output = strings.ReplaceAll(output, "^|^", `"`)
	wrappedOut := "[" + output + "]"

	err = json.Unmarshal([]byte(wrappedOut), &logs)

	if err != nil {
		fmt.Println("Cannot parse project", project, ":", err)
		return
	}

	i := 0 // output index
	for _, log := range logs {
		if !filter(config.Filters, log) {
			continue
		}

		logs[i] = log
		logs[i].Project = project

		msg := log.Message

		if !config.FullMessage && len(msg) > MAX_COMMIT_LEN {
			logs[i].Message = msg[:MAX_COMMIT_LEN] + "..."
		}

		i++
	}

	logs = logs[:i]

	return
}
