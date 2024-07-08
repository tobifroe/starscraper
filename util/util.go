package util

import (
	"encoding/csv"
	"os"

	"github.com/tobifroe/starscraper/types"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func WriteToCSV(users []types.User) {
	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"email", "name", "login"}
	writer.Write(headers)
	for _, row := range users {
		s := []string{
			row.Email,
			row.Name,
			row.Login,
		}
		writer.Write(s)
	}
}

// TODO implement/document client secret parsing, Write to Docs mode
func WriteToGoogleDocs(sheetFlag *string, allUsers []types.User) {
	return
}
