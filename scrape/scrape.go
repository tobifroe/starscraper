package scrape

import (
	"context"
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

func Scrape(token string, repo string, owner string, output string, verbose bool) {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	variables := map[string]interface{}{
		"repo":   githubv4.String(repo),
		"owner":  githubv4.String(owner),
		"cursor": (*githubv4.String)(nil), // Null after argument to get first page.
	}

	if token == "" {
		token = os.Getenv("GH_TOKEN")
	}

	if token != "" {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
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
				if v.Node.Email != "" && verbose {
					fmt.Printf("%s (%s) - %s\n", v.Node.Name, v.Node.Login, v.Node.Email)
				}
			}
			if !query.Repository.Stargazers.PageInfo.HasNextPage {
				break
			}
			variables["cursor"] = githubv4.NewString(query.Repository.Stargazers.PageInfo.EndCursor)
		}

		util.WriteToCSV(allUsers, output)
		s.Stop()
		fmt.Println("Success.")
		fmt.Printf("Wrote stargazer data to %s \n", output)

	} else {
		fmt.Println("No Github token supplied. Either pass the -token flag, set up a .env file or set the GH_TOKEN environment variable.")
	}
}
