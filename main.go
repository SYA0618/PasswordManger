package main

import (
	"log"
	"passwordManger/cmd"
	"passwordManger/internal/storage"
	"passwordManger/internal/utils"

	"github.com/spf13/cobra"
)

func main() {
	if err := storage.InitDB(); err != nil {
		log.Fatal("❌ Failed to initialize database:", err)
	}
	defer storage.Close()
	rootCmd := &cobra.Command{
		Use:   "passmanger",
		Short: "A simple CLI password manager",
		Long:  `A simple password manager that allows you to store and retrieve passwords securely.`,
	}

	rootCmd.AddCommand(cmd.AddCmd)
	utils.CheckErr(rootCmd.Execute(), "Failed to execute root command")

}
