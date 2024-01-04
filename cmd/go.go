/*
Copyright © 2024 SAM MIDGLEY sam@midgley.dev
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "A \"fun\" tic tac toe cli game",
	Long: `A \"fun\" tic tac toe cli game!
			Play with a friend or against the unbeatable AI, how fun.`,
	Run: func(cmd *cobra.Command, args []string) {
		ai, _ := cmd.Flags().GetBool("ai")

		play(ai)
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
	goCmd.Flags().Bool("ai", false, "Use the unbeatable AI")
}
