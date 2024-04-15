package aggregator

import "regexp"

// Used custom chars for json in order to allow escaping and then marshaling the string
var commitFormat = `--pretty=format:{
	^|^hash^|^: ^|^%h^|^,
	^|^date^|^: ^|^%ad^|^,
	^|^subject^|^: ^|^%s^|^,
	^|^author^|^: {
		^|^name^|^: ^|^%aN^|^,
		^|^email^|^: ^|^%aE^|^
	}
},`

var trailingComma = regexp.MustCompile(",$")

var MAX_COMMIT_LEN = 75
