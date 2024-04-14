package aggregator

import (
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/panjf2000/ants/v2"
	"github.com/sebcej/githis/utils"
)

func GetLogs(sources []Source, config Config, extraArgs []string) (logs []Log) {
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
			wg.Done()
		}
	}()

	for _, source := range sources {
		indexFolder(source, rc, &wg, config, extraArgs)
	}

	defer ants.Release()

	wg.Wait()
	close(rc)

	fmt.Println("Total logs: ", len(logs), "\n")

	sort.SliceStable(logs, func(i, j int) bool {
		prevTime, _ := utils.ParseLogDate(logs[i].Date)
		nextTime, _ := utils.ParseLogDate(logs[j].Date)

		return prevTime.After(nextTime)
	})

	if len(logs) > config.Filters.Limit {
		logs = logs[:config.Filters.Limit]
	}

	return
}

func indexFolder(source Source, rc chan []Log, wg *sync.WaitGroup, config Config, extraArgs []string) {
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
				projectLogs := getLogsFromGit(project.Name(), path, config, extraArgs)

				rc <- projectLogs
			})
		}
	}

}
