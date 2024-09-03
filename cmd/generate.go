package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"

	"time"

	"github.com/spf13/cobra"
)

var (
	startDate  string
	endDate    string
	maxCommits int
)

const (
	filePath = "README.md"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate fake git history",
	Long:  `Generate a fake git history for the specified repository within the given date range.`,
	Run: func(cmd *cobra.Command, args []string) {
		// args here is positional args
		start, err := parseDate(startDate)
		if err != nil {
			log.Fatalf("Failed to parse start date: %v", err)
		}

		end, err := parseDate(endDate)

		if err != nil {
			log.Fatalf("Failed to parse end date: %v", err)
			return
		}

		if start.After(end) {
			log.Fatalf("Start date is after end date")
			return
		}

		generateCommits(start, end, maxCommits)
	},
}

func parseDate(date string) (time.Time, error) {
	return time.Parse(time.DateOnly, date)
}

func generateCommits(startDate, endDate time.Time, maxCommits int) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < maxCommits; i++ {
		commitDate := startDate.AddDate(0, 0, r.Intn(int(endDate.Sub(startDate).Hours()/24)))
		commitMessage := randomCommitMessage()

		// fileContent := fmt.Sprintf("Commit %d: %s on %s\n", i+1, commitMessage, commitDate.Format(time.DateOnly))
		// err := os.WriteFile(filePath, []byte(fileContent), 0666)
		// if err != nil {
		// 	log.Fatalf("Failed to write to file: %v", err)
		// 	return
		// }
		commitDateStr := commitDate.Format(time.DateOnly)

		gitCmd := fmt.Sprintf("git commit --allow-empty -m \"%s\" --date \"%s\"", commitMessage, commitDateStr)
		cmd := exec.Command("bash", "-c", gitCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to run git command: %v", err)
			return
		}
	}

}

func randomCommitMessage() string {
	commitMessages := []string{
		"Fix bug in module X",
		"Update documentation",
		"Add new feature Y",
		"Refactor code for better performance",
		"Improve error handling",
		"Optimize algorithm Z",
		"Remove deprecated functions",
		"Add tests for feature A",
		"Update dependencies",
		"Code cleanup and formatting",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return commitMessages[r.Intn(len(commitMessages))]
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&startDate, "start", "s", "", "Start date (YYYY-MM-DD)")
	generateCmd.Flags().StringVarP(&endDate, "end", "e", "", "End date (YYYY-MM-DD)")
	generateCmd.Flags().IntVarP(&maxCommits, "commits", "m", 5, "Maximum number of commits per day")

	generateCmd.MarkFlagRequired("start")
	generateCmd.MarkFlagRequired("end")
	generateCmd.MarkFlagRequired("commits")
}
