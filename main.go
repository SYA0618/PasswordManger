package main

import (
	"passwordManger/cmd"
	"passwordManger/internal/utils"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "passmanger",
		Short: "A simple CLI password manager",
		Long:  `A simple password manager that allows you to store and retrieve passwords securely.`,
	}

	rootCmd.AddCommand(cmd.AddCmd)
	utils.CheckErr(rootCmd.Execute(), "Failed to execute root command")

}
