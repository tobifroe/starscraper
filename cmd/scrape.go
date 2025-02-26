package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tobifroe/starscraper/scrape"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrapes stargazer data, writes to output.csv",
	Run: func(cmd *cobra.Command, args []string) {
		owner := cmd.Flag("owner").Value.String()
		repo := cmd.Flag("repo").Value.String()
		token := cmd.Flag("token").Value.String()
		output := cmd.Flag("output").Value.String()
		verbose, _ := cmd.Flags().GetBool("verbose")
		scrape.Scrape(token, repo, owner, output, verbose)
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	scrapeCmd.Flags().String("repo", "", "Repository to scrape")
	scrapeCmd.Flags().String("owner", "", "Repository owner")
	scrapeCmd.Flags().String("token", "", "Github PAT")
	scrapeCmd.Flags().String("output", "output.csv", "Output file")
	scrapeCmd.Flags().BoolP("verbose", "v", false, "Verbose output")

	err := scrapeCmd.MarkFlagRequired("repo")
	if err != nil {
		panic(err)
	}
	err = scrapeCmd.MarkFlagRequired("owner")
	if err != nil {
		panic(err)
	}
}
