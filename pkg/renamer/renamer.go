package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"CineLabeler/pkg/searcher/omdb"
	"CineLabeler/pkg/searcher/tmdb"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func extractTitleAndYear(fileName, extension string) (string, string, error) {
	patterns := []regexp.Regexp{
		*regexp.MustCompile(`([^.]+)\.([0-9]{4})\..*(1080p|720p|WEBRip)\..*(WEB|WEB-DL|BluRay|x264|H264|HEVC|EAC3).*` + regexp.QuoteMeta(extension)),
		*regexp.MustCompile(`(.+)\.\((\d{4})\)` + regexp.QuoteMeta(extension)),
		*regexp.MustCompile(`([^.]+)\.(\d{4})` + regexp.QuoteMeta(extension)),
		*regexp.MustCompile(`([^.]+)` + regexp.QuoteMeta(extension)),
	}

	for _, pattern := range patterns {
		matches := pattern.FindStringSubmatch(fileName)
		if matches != nil {
			title := strings.NewReplacer(".", " ", "_", " ").Replace(matches[1])
			year := ""
			if len(matches) > 2 {
				year = matches[2]
			}
			return title, year, nil
		}
	}
	return "", "", fmt.Errorf("file does not match expected patterns, skipped renaming: %s", fileName)
}

func RenameFile(originalPath, destinationDirectory string, varEnv map[string]string) (string, error) {
	directory, originalName := filepath.Split(originalPath)

	extension := filepath.Ext(originalName)
	title, year, err := extractTitleAndYear(originalName, extension)
	if err != nil {
		return "", err
	}

	if destinationDirectory != "" {
		directory = destinationDirectory
	}

	return processFileRename(directory, title, year, originalPath, extension, varEnv)
}

func processFileRename(directory, title, year, originalPath, extension string, varEnv map[string]string) (string, error) {
	newTitle, newYear, err := searchTitle(title, year, varEnv)
	if err != nil {
		return "", err
	}

	newFileName := fmt.Sprintf("%s.(%s)%s", formatTitle(newTitle), newYear, extension)

	newFilePath := filepath.Join(directory, newFileName)

	if err := os.Rename(originalPath, newFilePath); err != nil {
		return "", err
	}
	return fmt.Sprintf("File renamed to: %s", newFileName), nil
}

func searchTitle(title, year string, varEnv map[string]string) (string, string, error) {
	if tmdbApiKey, ok := varEnv["tmdb"]; ok && tmdbApiKey != "" {
		if data, err := tmdb.SearchTMDB(title, year, tmdbApiKey, varEnv["lang"]); err == nil && data["title"] != "" {
			return data["title"], data["year"], nil
		}
	}
	if omdbApiKey, ok := varEnv["omdb"]; ok && omdbApiKey != "" {
		if data, err := omdb.SearchOMDB(title, year, omdbApiKey); err == nil && data["title"] != "" {
			return data["title"], data["year"], nil
		}
	}
	return "", "", fmt.Errorf("title not found for: %s (%s)", title, year)
}

func formatTitle(title string) string {
	words := strings.Split(title, " ")
	caser := cases.Title(language.Und)
	for i, word := range words {
		words[i] = caser.String(word)
	}
	return strings.Join(words, ".")
}
