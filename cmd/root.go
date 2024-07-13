package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "starscraper",
	Short: "Github Stargazer data getter.",
	Long:  `Starscraper is a simple application that returns public information for the stargazers of a given Github repo.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
