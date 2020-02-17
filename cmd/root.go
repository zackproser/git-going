package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	projectName string
	projectSlug string
)

var rootCmd = &cobra.Command{
	Use: "gitgoing",
	Short: "Git-going is a Github repository scaffolding tool",
	Long: "Creates and configures new local and remote repositories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gitgoing command running...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&projectName, "name", "n",  "", "The human legible name for your new project (Used in titles, etc)")
	rootCmd.PersistentFlags().StringVarP(&projectSlug, "slug", "s", "", "project slug (for URLs / Directory names)")

	rootCmd.AddCommand()
}