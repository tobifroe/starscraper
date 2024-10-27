package util

import (
	"bytes"
	"encoding/csv"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tobifroe/starscraper/types"
)

const outfile = "output.csv"

func TestWriteToCSV(t *testing.T) {
	// Mock data for testing
	users := []types.User{
		{Email: "user1@example.com", Name: "User One", Login: "user1"},
		{Email: "user2@example.com", Name: "User Two", Login: "user2"},
	}

	// Execute the function
	WriteToCSV(users, outfile)

	// Read the written file
	file, err := os.Open(outfile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	// Check if the CSV contents match the expected data
	expected := [][]string{
		{"email", "name", "login"},
		{"user1@example.com", "User One", "user1"},
		{"user2@example.com", "User Two", "user2"},
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("CSV content mismatch.\nExpected: %v\nGot: %v", expected, lines)
	}

	// Validate CSV format
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(file)
	csvContent := buf.String()

	// Checking for CSV format correctness using regular expression
	expectedCSVFormat := "^email,name,login\n[^,\n]+,[^,\n]+,[^,\n]+\n[^,\n]+,[^,\n]+,[^,\n]+\n$"
	matched, err := regexp.MatchString(expectedCSVFormat, csvContent)
	if err != nil {
		t.Fatal(err)
	}

	if !matched {
		t.Error("CSV format doesn't match the expected structure")
	}
}

func TestContains(t *testing.T) {
	matches := "foo"
	noMatch := "bar"
	testArray := []string{"foo", "foobizz", "buzz"}
	assert.True(t, Contains(testArray, matches))
	assert.False(t, Contains(testArray, noMatch))
}

func TestMain(m *testing.M) {
	// Run tests
	exitVal := m.Run()

	// Clean up after tests
	err := os.Remove(outfile)
	if err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
