package aggregator

import "regexp"

// Used custom chars for json in order to allow escaping and then marshaling the string
var commitFormat = `--pretty=format:{
	^|^hash^|^: ^|^%h^|^,
	^|^date^|^: ^|^%ad^|^,
	^|^branch^|^: ^|^%D^|^,
	^|^subject^|^: ^|^%s^|^,
	^|^author^|^: {
		^|^name^|^: ^|^%aN^|^,
		^|^email^|^: ^|^%aE^|^
	}
},`

var trailingComma = regexp.MustCompile(",$")
var matchRepoEnd = regexp.MustCompile(".git$")

var matchDatePartialDay = regexp.MustCompile(`^\d\d$`)
var matchDatePartialDayMonth = regexp.MustCompile(`^(\d\d)-(\d\d)$`)

var matchGithubRepoStart = regexp.MustCompile("^.+github.com(:|/)")
var matchGitlabRepoStart = regexp.MustCompile("^.+gitlab.com(:|/)")

var MAX_COMMIT_LEN = 75
