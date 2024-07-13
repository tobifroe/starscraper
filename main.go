package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"github.com/tobifroe/starscraper/types"
	"github.com/tobifroe/starscraper/util"
	"golang.org/x/oauth2"
)

var query struct {
	Repository struct {
		Description string
		Stargazers  struct {
			TotalCount int
			PageInfo   struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
			Edges []struct {
				Node struct {
					Email string
					Name  string
					Login string
				}
			}
		} `graphql:"stargazers(first: 100, after: $cursor)"`
	} `graphql:"repository(owner: $owner, name: $repo)"`
}

func main() {

	// init Flags
	tokenFlag := flag.String("token", "", "Github Token")
	repoFlag := flag.String("repo", "", "Github Repo")
	ownerFlag := flag.String("owner", "", "Github Repo Owner")
	outputFlag := flag.String("output", "output.csv", "Output file name")
	debugFlag := flag.Bool("debug", false, "Verbose output for debugging")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	variables := map[string]interface{}{
		"repo":   githubv4.String(*repoFlag),
		"owner":  githubv4.String(*ownerFlag),
		"cursor": (*githubv4.String)(nil), // Null after argument to get first page.
	}

	if *tokenFlag == "" {
		*tokenFlag = os.Getenv("GH_TOKEN")
	}

	if *repoFlag == "" {
		fmt.Println("Please specify a repository.")
		return
	}

	if *ownerFlag == "" {
		fmt.Println("Please specify a repository owner.")
		return
	}

	if *tokenFlag != "" {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: *tokenFlag},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := githubv4.NewClient(tc)

		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		fmt.Println("Getting stargazers...")
		s.Start()

		var allUsers []types.User
		for {
			err := client.Query(ctx, &query, variables)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, v := range query.Repository.Stargazers.Edges {
				if v.Node.Email != "" {
					allUsers = append(allUsers, types.User{
						Email: v.Node.Email,
						Name:  v.Node.Name,
						Login: v.Node.Login,
					})
				}
				if v.Node.Email != "" && *debugFlag {
					fmt.Printf("%s (%s) - %s\n", v.Node.Name, v.Node.Login, v.Node.Email)
				}
			}
			if !query.Repository.Stargazers.PageInfo.HasNextPage {
				break
			}
			variables["cursor"] = githubv4.NewString(query.Repository.Stargazers.PageInfo.EndCursor)
		}

		util.WriteToCSV(allUsers, *outputFlag)
		s.Stop()
		fmt.Println("Success.")
		fmt.Printf("Wrote stargazer data to %s \n", *outputFlag)

	} else {
		fmt.Println("No Github token supplied. Either pass the -token flag, set up a .env file or set the GH_TOKEN environment variable.")
	}
}
