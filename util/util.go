package util

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/tobifroe/starscraper/types"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
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
			row.Login,
			row.Name,
		}
		writer.Write(s)
	}
}

func WriteToGoogleDocs(sheetFlag *string, allUsers []types.User) {
	data, err := os.ReadFile("client-secret.json")
	if err != nil {
		fmt.Println("Unable to read client secret file:", err)
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		fmt.Println("Unable to parse client secret file to config:", err)
	}

	sheetClient := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(sheetClient)

	sheetID := *sheetFlag

	usersSheet, err := service.FetchSpreadsheet(sheetID)
	if err != nil {
		fmt.Println("Unable to fetch sheet:", err)
	}

	sheet, err := usersSheet.SheetByIndex(0)

	var flatUsernames []string

	if len(sheet.Columns) > 1 {
		for _, v := range sheet.Columns[2] {
			if v.Value != "" {
				flatUsernames = append(flatUsernames, v.Value)
			}
		}

		insert := len(sheet.Columns[2])
		fmt.Println("Sheet has data")

		for _, v := range allUsers {
			if Contains(flatUsernames, v.Login) {
				fmt.Printf("%s (%s) - %s\n", v.Name, v.Login, v.Email)
				continue
			}
			sheet.Update(insert, 1, v.Name)
			sheet.Update(insert, 2, v.Login)
			sheet.Update(insert, 3, v.Email)
			insert++
		}
	} else {
		fmt.Println("Sheet is empty")
		for i, v := range allUsers {
			if err != nil {
				fmt.Println(err)
				break
			}
			sheet.Update(i+1, 1, v.Name)
			sheet.Update(i+1, 2, v.Login)
			sheet.Update(i+1, 3, v.Email)
		}
	}
	syncErr := sheet.Synchronize()
	if syncErr != nil {
		fmt.Println(err)
	}
}
