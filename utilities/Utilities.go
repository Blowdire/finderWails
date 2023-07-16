package utilities

type SearchResult struct {
	Filename string
	Filepath string
	IsDir    bool
}

type ListingResults struct {
	Files       []SearchResult
	Directories []SearchResult
}
