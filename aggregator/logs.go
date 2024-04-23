package aggregator

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func getLogsFromGit(project, dir string, config Config, extraArgs []string) (logs []Log) {
	builtArgs := []string{"log", "--all", "--decorate", "--date=format:%Y-%m-%d %H:%M:%S", commitFormat}
	builtArgs = append(builtArgs, extraArgs...)

	if config.Pull {
		pullCmd := exec.Command("git", "pull", "--quiet")
		pullCmd.Dir = dir

		_, err := pullCmd.CombinedOutput()
		if err != nil {
			return
		}
	}

	cmdLogs := exec.Command("git", builtArgs...)
	cmdLogs.Dir = dir

	out, err := cmdLogs.CombinedOutput()
	if err != nil {
		return
	}

	urlBase := getUrlFromProject(dir)

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
		logs[i].CommitUrl = urlBase.getCommitUrl(log.Hash)

		msg := log.Message

		if !config.FullMessage && len(msg) > MAX_COMMIT_LEN {
			logs[i].Message = msg[:MAX_COMMIT_LEN] + "..."
		}

		i++
	}

	logs = logs[:i]

	return
}

func getUrlFromProject(dir string) commitBase {
	cmdGetRemotes := exec.Command("git", "ls-remote", "--get-url")
	cmdGetRemotes.Dir = dir

	rawUrls, err := cmdGetRemotes.CombinedOutput()
	if err != nil {
		return ""
	}

	strUrls := string(rawUrls)
	urls := strings.Split(strUrls, "\n")

	if len(urls) == 0 {
		return ""
	}

	url := urls[0]

	return commitBase(url)
}

func (typeUrl commitBase) getCommitUrl(commit string) string {
	url := string(typeUrl)

	if strings.Contains(url, "github.com") {
		url = matchRepoEnd.ReplaceAllString(url, "")
		url = matchGithubRepoStart.ReplaceAllString(url, "")

		url = "https://github.com/" + url + "/commit/" + commit

		return url
	}

	if strings.Contains(url, "gitlab.com") {
		url = matchRepoEnd.ReplaceAllString(url, "")
		url = matchGitlabRepoStart.ReplaceAllString(url, "")

		url = "https://gitlab.com/" + url + "/-/commit/" + commit

		return url
	}

	return ""
}
