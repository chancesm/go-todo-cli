package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
	},
}
