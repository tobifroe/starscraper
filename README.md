# Starscraper

[![Coverage Status](https://coveralls.io/repos/github/tobifroe/starscraper/badge.svg?branch=main)](https://coveralls.io/github/tobifroe/starscraper?branch=main)

Starscraper is a simple application that returns public information for the stargazers of a given Github repo.

## Installation

TODO

## Usage

Starscraper needs a GitHub PAT to operate. It needs to have the `repo:public_repo` and `user:email` scopes.
PATs can be provided to Starscraper by:

- Setting GH_PAT environment variable
- Creating a .env file in the directory Starscraper is run from. Check `.env.example`.
- Passing the --token flag

```
starscraper scrape [flags]
```

### Options

```
  -h, --help            help for scrape
      --output string   Output file (default "output.csv")
      --owner string    Repository owner
      --repo string     Repository to scrape
      --token string    Github PAT
  -v, --verbose         Verbose output
```

## Contributing

TODO
