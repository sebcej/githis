package aggregator

type Source struct {
	Name string
	Path string
}
type Config struct {
	Filters     Filters
	FullMessage bool // True if commit message needs to be full
	Raw         bool // Show RAW json git output

	Offset  int // Days of offset, defaults to 0 (today)
	FromDay string
	ToDay   string
}

type Filters struct{}

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
