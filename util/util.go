package util

import (
	"encoding/csv"
	"os"

	"github.com/tobifroe/starscraper/types"
)

func WriteToCSV(users []types.User, output string) {
	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	//nolint:errcheck
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"email", "name", "login"}
	err = writer.Write(headers)
	if err != nil {
		panic(err)
	}
	for _, row := range users {
		s := []string{
			row.Email,
			row.Name,
			row.Login,
		}
		err = writer.Write(s)
		if err != nil {
			panic(err)
		}
	}
}

func WriteToGoogleDocs(sheetFlag *string, allUsers []types.User) {
	// TODO implement/document client secret parsing, Write to Docs mode
}
