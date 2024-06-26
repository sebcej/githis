package aggregator

type Source struct {
	Name string
	Path string
}
type Config struct {
	Filters     Filters
	FullMessage bool // True if commit message needs to be full
	Raw         bool // Show RAW json git output
	Pull        bool // Pull automatically before parsing logs
	Reverse     bool // Reverse the order of logs
}

type Filters struct {
	Offset  int      // Days of offset, defaults to 0 (today)
	Authors []string // Authors of commit
	From    string
	To      string
	Day     string // Automatically set from and to to the selected day
	Limit   int
}

type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Log struct {
	Date    string       `json:"date"`
	Hash    string       `json:"hash"`
	Author  CommitAuthor `json:"author"`
	Message string       `json:"subject"`
	Branch  string       `json:"branch"`

	Project   string `json:"project"`
	CommitUrl string `json:"commitUrl"`
}

type commitBase string
