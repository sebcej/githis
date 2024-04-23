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
}

type Filters struct {
	Offset  int      // Days of offset, defaults to 0 (today)
	Authors []string // Authors of commit
	FromDay string
	ToDay   string
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
	Project string
}
