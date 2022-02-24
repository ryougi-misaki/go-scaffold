package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "生成一个a.txt",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := os.Create("a.txt")
		return err
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)
}
