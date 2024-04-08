package aggregator

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func GetLogs(sources []Source, filters Filters, extraArgs []string) (logs []Log) {
	rc := make(chan []Log)
	threads := 0

	for _, source := range sources {
		fmt.Println("Source", source)

		threads++
		go indexFolder(source, rc, filters, extraArgs)
	}

	// Parallelize folders scanning
	for {
		logsPart := <-rc

		logs = append(logs, logsPart...)

		threads--

		if threads == 0 {
			break
		}
	}

	fmt.Println("Available logs: ", len(logs))

	return
}

func indexFolder(source Source, rc chan []Log, filters Filters, extraArgs []string) {
	projects, err := os.ReadDir(source.Path)
	if err != nil {
		rc <- []Log{}
		return
	}

	var logs []Log

	for _, project := range projects {
		path := source.Path + "/" + project.Name()
		_, err := os.Stat(path + "/.git")

		if project.IsDir() && !os.IsNotExist(err) {
			projectLogs := getFromGit(path, filters, extraArgs)
			logs = append(logs, projectLogs...)
		}
	}

	rc <- logs
}

func getFromGit(dir string, filters Filters, extraArgs []string) (logs []Log) {
	builtArgs := []string{"log", "--date=format:%Y-%m-%d %H:%M:%S", commitFormat}
	builtArgs = append(builtArgs, extraArgs...)

	cmd := exec.Command("git", builtArgs...)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err")
	}

	wrappedOut := "[" + string(out) + "]"

	err = json.Unmarshal([]byte(wrappedOut), &logs)
	if err != nil {
		return
	}

	return
}
