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

var MAX_COMMIT_LEN = 75
