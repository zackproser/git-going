package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	projectName string
	projectSlug string
	log         = logrus.New()
)

var rootCmd = &cobra.Command{
	Use:   "gitgoing",
	Short: "Git-going is a Github repository scaffolding tool",
	Long:  "Creates and configures new projects",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("gitgoing command running...")
	},
}

// Execute is the entrypoint to the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)

	rootCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "The human legible name for your new project (Used in titles, etc)")
	rootCmd.PersistentFlags().StringVarP(&projectSlug, "slug", "s", "", "project slug (for URLs / Directory names)")

	rootCmd.AddCommand()
}
