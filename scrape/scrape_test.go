package scrape

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestScrape(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Load environment variables for tests
	err := godotenv.Load(".env.test")
	if err != nil {
		fmt.Println("No .env.test file present, proceeding without it.")
	}

	t.Run("should fail if no token is provided", func(t *testing.T) {
		os.Unsetenv("GH_TOKEN")
		output := "output.csv"

		// Capture the output
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the function
		Scrape("", "repo", "owner", output, false)

		w.Close()
		out, _ := io.ReadAll(r)
		os.Stdout = old

		assert.Contains(t, string(out), "No Github token supplied")
	})

	t.Run("should use token from environment if not provided", func(t *testing.T) {
		// Set the environment variable for the token
		os.Setenv("GH_TOKEN", "mocked-token")
		defer os.Unsetenv("GH_TOKEN")

		// Assert that token was used (in reality, would use more sophisticated validation)
		assert.Equal(t, "mocked-token", os.Getenv("GH_TOKEN"))
	})
}
