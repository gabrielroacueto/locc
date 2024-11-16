package fsparser

import (
	"os"
	"path/filepath"
)

// GetDirectoryContents returns a slice of strings containing the names of files
// and directories in the specified directory path. It returns an error if the
// directory cannot be read or if the path is invalid.
func GetDirectoryContents(directory string) ([]string, error) {
	// Expand the directory path if it contains ~
	if len(directory) > 0 && directory[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		directory = filepath.Join(homeDir, directory[1:])
	}

	// Clean the path to handle . or ..
	directory = filepath.Clean(directory)

	// Open the directory
	dir, err := os.Open(directory)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Read directory contents
	entries, err := dir.ReadDir(-1) // -1 means read all entries
	if err != nil {
		return nil, err
	}

	// Create slice to store names
	contents := make([]string, 0, len(entries))

	// Add each entry name to the slice
	for _, entry := range entries {
		contents = append(contents, entry.Name())
	}

	return contents, nil
}
