package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/BaseTechStack/bux/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "bux",
	Aliases: []string{"basenuxt"},
	Short:   "BaseUX CLI - A modern Nuxt scaffolding framework",
	Long: `BaseUX (bux) CLI is a powerful tool for building modern Nuxt applications.
It provides scaffolding, code generation, and development utilities. Works best with the BaseNuxt framework.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Skip version check for version and upgrade commands
		if cmd.Name() != "version" && cmd.Name() != "upgrade" {
			if release, err := version.CheckLatestVersion(); err == nil {
				info := version.GetBuildInfo()
				latestVersion := strings.TrimPrefix(release.TagName, "v")
				// Only show update message if there's actually an update
				if version.HasUpdate(info.Version, latestVersion) {
					fmt.Print(version.FormatUpdateMessage(
						info.Version,
						latestVersion,
						release.HTMLURL,
						release.Body,
					))
				} else {
					fmt.Printf("\n✨ You're up to date! Using the latest version %s\n", info.Version)
				}
			} else {
				fmt.Println("Failed to check for updates:", err)
			}
		}
	},
}

func Execute() error {
	// Check if the command was invoked as basenuxt (deprecated)
	if len(os.Args) > 0 && strings.Contains(os.Args[0], "basenuxt") {
		fmt.Println("⚠️  Warning: The 'basenuxt' command is deprecated and will be removed in a future version.")
		fmt.Println("   Please use 'bux' instead, which provides the same functionality.")
	}
	return rootCmd.Execute()
}

func init() {
	// Commands will be added here in subsequent steps
}
