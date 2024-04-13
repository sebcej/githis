package aggregator

import (
	"os"
	"sort"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
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

	sort.SliceStable(logs, func(i, j int) bool {
		prevTime, _ := time.Parse("2006-01-02 15:04:05", logs[i].Date)
		nextTime, _ := time.Parse("2006-01-02 15:04:05", logs[j].Date)

		return prevTime.After(nextTime)
	})

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
