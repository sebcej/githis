package aggregator

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func GetLogs(sources []Source, filters Filters, extraArgs []string) (logs []Log) {
	rc := make(chan []Log)
	var wg sync.WaitGroup

	// Parallelize folders scanning
	go func() {
		for {
			logsPart, more := <-rc

			if !more {
				return
			}

			logs = append(logs, logsPart...)
		}
	}()

	for _, source := range sources {
		indexFolder(source, rc, &wg, filters, extraArgs)
	}

	defer ants.Release()

	wg.Wait()
	close(rc)

	// Wait for chan to close
	time.Sleep(1 * time.Millisecond)

	sort.SliceStable(logs, func(i, j int) bool {
		prevTime, _ := time.Parse("2006-01-02 15:04:05", logs[i].Date)
		nextTime, _ := time.Parse("2006-01-02 15:04:05", logs[j].Date)

		return prevTime.Before(nextTime)
	})

	return
}

func indexFolder(source Source, rc chan []Log, wg *sync.WaitGroup, filters Filters, extraArgs []string) {
	projects, err := os.ReadDir(source.Path)
	if err != nil {
		rc <- []Log{}
		return
	}

	for _, project := range projects {
		path := source.Path + "/" + project.Name()
		_, err := os.Stat(path + "/.git")

		if project.IsDir() && !os.IsNotExist(err) {
			wg.Add(1)

			ants.Submit(func() {
				projectLogs := getFromGit(project.Name(), path, filters, extraArgs)

				rc <- projectLogs

				wg.Done()
			})
		}
	}

}

func getFromGit(project, dir string, filters Filters, extraArgs []string) (logs []Log) {
	builtArgs := []string{"log", "--all", "--date=format:%Y-%m-%d %H:%M:%S", commitFormat}
	builtArgs = append(builtArgs, extraArgs...)

	cmd := exec.Command("git", builtArgs...)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	output := string(out)
	output = trailingComma.ReplaceAllString(output, "")
	wrappedOut := "[" + output + "]"

	err = json.Unmarshal([]byte(wrappedOut), &logs)

	i := 0 // output index
	for _, log := range logs {
		if !filter(filters, log) {
			continue
		}
		logs[i] = log
		i++
	}

	logs = logs[:i]

	if err != nil {
		fmt.Println("git parse error", err)
		return
	}

	return
}
