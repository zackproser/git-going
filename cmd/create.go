package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	scaffold "github.com/zackproser/git-going/internal/scaffold"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	github "github.com/zackproser/git-going/internal/github"
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
	Use:   "create",
	Short: "Create a new project",
	Long:  "Creates local and remote git repos, scaffolds files, runs git config",
	Run: func(cmd *cobra.Command, args []string) {
		if projectName == "" {
			fmt.Println("You must provide a unique project name to create a new project")
			os.Exit(1)
		}

		// If author name was not supplied, attempt to look it up via git config
		if authorName == "" {
			userName, getUserErr := scaffold.GetGitUserName(projectSlug)
			if getUserErr != nil || userName == "" {
				log.WithFields(logrus.Fields{
					"Error": getUserErr,
				}).Debug("Error looking up username via git config")
				authorName = "Unknown"
			}
			authorName = userName
		}

		// If user doesn't provide a project slug, convert the project name
		// to a suitable slug
		if projectSlug == "" {
			projectSlug = convertProjectNameToSlug(projectName)
			log.Debug(fmt.Sprintf("Converted %s to slug: %s", projectName, projectSlug))
			scaffoldErr := scaffold.Create(projectName, projectSlug, authorName, log)
			if scaffoldErr != nil {
				log.WithFields(logrus.Fields{
					"Error": scaffoldErr,
				}).Debug("Error scaffolding new project")
				return
			}

			repository, remoteRepoCreateErr := github.CreateRepo(projectSlug)

			if remoteRepoCreateErr != nil {
				log.WithFields(logrus.Fields{
					"Error": remoteRepoCreateErr,
				}).Debug("Error creating remote repository")
				return
			}

			// Get the SSH URL for setting the local repository's remote URL
			if repository.GetSSHURL() != "" {
				originAddErr := scaffold.AddRemoteOrigin(projectSlug, repository.GetSSHURL())
				if originAddErr != nil {
					log.WithFields(logrus.Fields{
						"Error": originAddErr,
					}).Debug("Error adding remote origin")
				}
			} else {
				log.WithFields(logrus.Fields{
					"Error": errors.New("Repository SSH URL is empty"),
				}).Debug("Error retrieving SSH URL from created remote repo")
			}

			// Push the local repository to the created remote
			pushErr := scaffold.PushRepo(projectSlug)
			if pushErr != nil {
				log.WithFields(logrus.Fields{
					"Error": pushErr,
				}).Debug("Error pushing local repository to remote")
			}
		}
	},
}
