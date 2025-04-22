package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// ReplaceInFile replaces all occurrences of a string in a file
func ReplaceInFile(filePath, oldString, newString string) error {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	// Replace the string
	newContent := strings.ReplaceAll(string(content), oldString, newString)

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filePath, err)
	}

	return nil
}

// ReplaceInAllFiles performs a find/replace in all text files in a directory (recursively)
func ReplaceInAllFiles(dirPath, oldString, newString string) error {
	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dirPath)
	}

	// List of binary file extensions to skip
	binaryExts := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
		".svg": true, ".ico": true, ".pdf": true, ".zip": true,
		".gz": true, ".tar": true, ".mp3": true, ".mp4": true,
		".woff": true, ".woff2": true, ".ttf": true, ".eot": true,
		".o": true, ".a": true, ".so": true, ".dylib": true, ".dll": true, ".exe": true,
	}

	// Walk through all files in the directory
	return filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Skip binary files based on extension
		ext := strings.ToLower(filepath.Ext(path))
		if binaryExts[ext] {
			return nil
		}

		// Skip node_modules directory and .git directory (but not .github)
		if strings.Contains(path, "node_modules") || strings.Contains(path, "/.git/") || strings.HasSuffix(path, "/.git") {
			return nil
		}

		// Perform the replacement
		return ReplaceInFile(path, oldString, newString)
	})
}
