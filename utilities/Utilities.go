package utilities

type SearchResult struct {
	Filename string
	Filepath string
}

type ListingResults struct {
	Files       []SearchResult
	Directories []SearchResult
}
