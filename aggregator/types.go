package aggregator

type Source struct {
	Name string
	Path string
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
	Message string       `json:"message"`
}
