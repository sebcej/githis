package aggregator

import "regexp"

var commitFormat = `--pretty=format:{
	"hash": "%h",
	"date": "%ad",
	"subject": "%s",
	"author": {
		"name": "%aN",
		"email": "%aE"
	}
},`

var trailingComma = regexp.MustCompile(",$")
var matchRepoEnd = regexp.MustCompile(".git$")

var matchGitRepoStart = regexp.MustCompile("^.+github.com(:|/)")

var MAX_COMMIT_LEN = 75
