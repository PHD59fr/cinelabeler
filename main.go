package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"CineLabeler/pkg/renamer"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 || len(args) > 2 {
		fmt.Printf("Usage: %s <original_file.mkv> [destination_directory]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	originalFilePath := args[0]
	varEnv := getVarEnv()

	var destinationDirectory string
	if len(args) == 2 {
		destinationDirectory = args[1]
	}

	result, err := renamer.RenameFile(originalFilePath, destinationDirectory, varEnv)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

func getVarEnv() map[string]string {
	return map[string]string{
		"omdb": os.Getenv("OMDB_API_KEY"),
		"tmdb": os.Getenv("TMDB_API_KEY"),
		"lang": os.Getenv("LANG"),
	}
}
