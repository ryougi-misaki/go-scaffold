package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "sca",
	Short: "sca 项目构建工具",
	Long:  "sca 项目构建工具",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
