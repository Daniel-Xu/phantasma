package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "phantasma",
	Short: "Generate fake git history for a repository",
	Long: `phantasma is a CLI tool that generates a fake git history
for a repository. It can be used to create a more impressive-looking
contribution graph or to simulate long-term project activity.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
