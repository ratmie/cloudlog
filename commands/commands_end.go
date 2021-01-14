package commands

import (
	"github.com/spf13/cobra"
)

var (
	xxxCmd = &cobra.Command{
		Use: "xxx",
		Run: xxxCommand,
	}
)

func xxxCommand(cmd *cobra.Command, args []string) {
	if err := xxxAction(); err != nil {
		Exit(err, 1)
	}
}

func xxxAction() (err error) {
	// 実行したい内容
}

func init() {
	RootCmd.AddCommand(xxxCmd)
}
