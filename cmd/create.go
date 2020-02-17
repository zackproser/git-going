package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	scaffold "github.com/zackproser/git-going/pkg"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

func convertProjectNameToSlug(n string) string {
	var ret string
	spaces := regexp.MustCompile(`\s+`)
	ret = strings.Trim(n, " ")
	// Remove any duplicate whitespaces
	ret = spaces.ReplaceAllString(ret, " ")
	// Replace all spaces with hyphens
	ret = strings.Replace(ret, " ", "-", -1)
	ret = strings.ToLower(ret)
	return ret
}

var createCmd = &cobra.Command{
	Use: "create",
	Short: "Create a new project",
	Long: "Creates local and remote git repos, scaffolds files, runs git config",
	Run: func(cmd *cobra.Command, args []string) {
		if projectName == "" {
			fmt.Println("You must provide a unique project name to create a new project")
			os.Exit(1)
		}
		// If user doesn't provide a project slug, convert the project name
		// to a suitable slug
		if projectSlug == "" {
			projectSlug = convertProjectNameToSlug(projectName)
			fmt.Println("Converted projectSlug: ", projectSlug)
			scaffold.Create(projectName, projectSlug)
		}
	},
}
