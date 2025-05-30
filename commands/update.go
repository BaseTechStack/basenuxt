package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update BaseNuxt Core to the latest version",
	Long:  `Update BaseNuxt Core to the latest version. This command will update the core directory of your BaseNuxt project to the latest version available on GitHub.`,
	Run:   updateBaseCore,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateBaseCore(cmd *cobra.Command, args []string) {
	fmt.Println("Updating BaseNuxt Core...")
	err := updateCore()
	if err != nil {
		fmt.Printf("Error updating BaseNuxt Core: %v\n", err)
		return
	}
	fmt.Println("BaseNuxt Core updated successfully.")
}

func updateCore() error {
	// Define the core repository URL
	coreRepoURL := "https://github.com/BaseTechStack/bux.git"

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "basenuxt-core-update")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Clone the core repository
	gitCmd := exec.Command("git", "clone", coreRepoURL, tempDir)
	if output, err := gitCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone core repository: %v\n%s", err, output)
	}

	// Path to the core directory in the current project
	projectCoreDir := filepath.Join(".", "core")

	// Backup the current core directory
	backupDir := projectCoreDir + ".bak"
	if err := os.Rename(projectCoreDir, backupDir); err != nil {
		return fmt.Errorf("failed to backup current core directory: %v", err)
	}

	// Exclude nuxt.config.ts and ./structures
	excludePaths := []string{
		"nuxt.config.ts",
		"structures",
	}

	// Copy core files from temp directory to the project
	tempCoreDir := filepath.Join(tempDir, "core")
	err = filepath.Walk(tempCoreDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip excluded paths
		relPath, err := filepath.Rel(tempCoreDir, path)
		if err != nil {
			return err
		}

		for _, exclude := range excludePaths {
			if strings.Contains(relPath, exclude) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		targetPath := filepath.Join(projectCoreDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(targetPath, data, 0644)
	})

	if err != nil {
		// If there's an error, attempt to restore the backup
		os.RemoveAll(projectCoreDir)
		os.Rename(backupDir, projectCoreDir)
		return fmt.Errorf("failed to copy core files: %v", err)
	}

	// Remove the backup directory
	os.RemoveAll(backupDir)

	fmt.Println("Core directory updated successfully.")
	return nil
}
