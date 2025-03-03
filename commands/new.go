package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BaseTechStack/basenuxt/utils"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [project_name]",
	Short: "Create a new project",
	Long:  `Create a new project by cloning the basenuxt repository and setting up the directory.`,
	Args:  cobra.ExactArgs(1),
	Run:   createNewProject,
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func createNewProject(cmd *cobra.Command, args []string) {
	projectName := args[0]
	archiveURL := "https://github.com/BaseTechStack/basenuxt-source/archive/main.zip"

	// Create the project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Printf("Error creating project directory: %v\n", err)
		return
	}

	// Download the archive
	resp, err := http.Get(archiveURL)
	if err != nil {
		fmt.Printf("Error downloading project template: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Create a temporary file to store the zip
	tmpZip, err := os.CreateTemp("", "basenuxt-project-*.zip")
	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		return
	}
	defer os.Remove(tmpZip.Name())

	// Copy the zip content to the temporary file
	_, err = io.Copy(tmpZip, resp.Body)
	if err != nil {
		fmt.Printf("Error saving project template: %v\n", err)
		return
	}
	tmpZip.Close()

	// Unzip the file
	err = utils.Unzip(tmpZip.Name(), projectName)
	if err != nil {
		fmt.Printf("Error extracting project template: %v\n", err)
		return
	}

	// Loop through each file/directory in the basenuxt-source-main directory and move it to the project root
	files, err := os.ReadDir(filepath.Join(projectName, "basenuxt-source-main"))
	if err != nil {
		fmt.Printf("Error reading extracted files: %v\n", err)
		return
	}

	for _, f := range files {
		oldPath := filepath.Join(projectName, "basenuxt-source-main", f.Name())
		newPath := filepath.Join(projectName, f.Name())
		if err := os.Rename(oldPath, newPath); err != nil {
			fmt.Printf("Error moving file/directory %s: %v\n", f.Name(), err)
			return
		}
	}

	// Clean up the temporary directory
	os.RemoveAll(filepath.Join(projectName, "basenuxt-source-main"))

	// Get the absolute path of the new project directory
	absPath, err := filepath.Abs(projectName)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		return
	}

	fmt.Printf("New project '%s' created successfully at %s\n", projectName, absPath)
	
	// Install dependencies
	fmt.Println("Installing dependencies...")
	
	// Check if bun is available
	bunCmd := exec.Command("bun", "--version")
	bunAvailable := bunCmd.Run() == nil
	
	// Check if yarn is available
	yarnCmd := exec.Command("yarn", "--version")
	yarnAvailable := yarnCmd.Run() == nil
	
	// Change to the project directory
	originalDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}
	defer os.Chdir(originalDir) // Change back to original directory when done
	
	if err := os.Chdir(absPath); err != nil {
		fmt.Printf("Error changing to project directory: %v\n", err)
		return
	}
	
	// Run bun, yarn or npm install based on availability (in that order of preference)
	var installCmd *exec.Cmd
	if bunAvailable {
		fmt.Println("Using bun to install dependencies...")
		installCmd = exec.Command("bun", "install")
	} else if yarnAvailable {
		fmt.Println("Using yarn to install dependencies...")
		installCmd = exec.Command("yarn", "install")
	} else {
		fmt.Println("Using npm to install dependencies...")
		installCmd = exec.Command("npm", "install")
	}
	
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	
	if err := installCmd.Run(); err != nil {
		fmt.Printf("Error installing dependencies: %v\n", err)
		fmt.Println("You may need to run 'npm install', 'yarn install', or 'bun install' manually.")
	} else {
		fmt.Println("Dependencies installed successfully!")
	}
	
	fmt.Println("You can now start working on your new project!")
}
