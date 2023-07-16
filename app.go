package main

import (
	"Finder/utilities"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

}

// Greet returns a greeting for the given name

func (a *App) GetPartitions() []string {
	r := []string{}
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			drivePath := string(drive) + ":\\"
			r = append(r, drivePath)
			f.Close()
		}
	}
	return r
}

func (a *App) ListDir(path string) utilities.ListingResults {
	fmt.Println("Listing directory:", path)
	files_listed := []utilities.SearchResult{}
	dirs_listed := []utilities.SearchResult{}
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return utilities.ListingResults{Files: files_listed, Directories: dirs_listed}
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())

		if file.IsDir() {

			dirs_listed = append(dirs_listed, utilities.SearchResult{Filename: file.Name(), Filepath: filePath, IsDir: true})
		} else {
			files_listed = append(files_listed, utilities.SearchResult{Filename: file.Name(), Filepath: filePath, IsDir: false})
		}
	}
	return utilities.ListingResults{Files: files_listed, Directories: dirs_listed}
}

func (a *App) Search(rootPath string, pattern string) []utilities.SearchResult {
	// Set the root directory for scanning

	var wg sync.WaitGroup
	var wgScan sync.WaitGroup                      // Create a wait group to synchronize goroutines
	filePaths := make(chan utilities.SearchResult) // Create a channel to receive file paths
	dirPaths := make(chan string)
	waitChan := make(chan struct{}, 9000)

	// Launch a goroutine to traverse the filesystem and send file paths to the channel
	wgScan.Add(1)
	go func() {

		traverse(rootPath, filePaths, &wgScan, waitChan)

		fmt.Println("Done traversing filesystem")

	}()
	filteredFiles := make(chan utilities.SearchResult)
	wg.Add(1)
	go func() {
		defer close(filteredFiles)

		processFiles(filePaths, pattern, filteredFiles, &wg)
	}()
	fmt.Println("Done processing files")
	var wgFinal sync.WaitGroup

	files := make([]utilities.SearchResult, 0)
	wgFinal.Add(1)
	go func() {
		defer wgFinal.Done()
		for file := range filteredFiles {
			files = append(files, file)
		}
	}()
	wgScan.Wait()
	close(filePaths)
	close(dirPaths)
	close(waitChan)
	wg.Wait()
	wgFinal.Wait()
	fmt.Println(len(files))
	return files
}

func traverse(dir string, filePaths chan<- utilities.SearchResult, wg *sync.WaitGroup, waitChan chan struct{}) {
	defer wg.Done()
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())

		if file.IsDir() {

			wg.Add(1)

			select {
			case waitChan <- struct{}{}:
				go func() {
					defer func() { <-waitChan }()
					traverse(filePath, filePaths, wg, waitChan)
				}()
			default:
				traverse(filePath, filePaths, wg, waitChan)
			}
			// Recurse into subdirectories
		} else {
			filePaths <- utilities.SearchResult{Filename: file.Name(), Filepath: filePath, IsDir: false}
			// Send file path to the channel
		}
	}
}
func (a *App) OpenFile(path string) {
	fmt.Println("Opening file:", path)
	exec.Command("explorer", "/select,", path).Run()
}

// processFiles is the function that handles the processing of file paths received from the channel
func processFiles(filePaths <-chan utilities.SearchResult, pattern string, filteredFileChannel chan<- utilities.SearchResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for filePath := range filePaths {
		//Process the file path (e.g., perform operations on the file)
		if fuzzy.Match(pattern, filePath.Filename) {
			filteredFileChannel <- filePath
		}
	}
}
