package cmd

import (
	"errors"
	"fmt"
	"passwordManger/internal/utils"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Replace with actual implementation
		fmt.Println("[add] This will add a new password entry (to be implemented)")
		err := errors.New("not implemented")
		utils.CheckErr(err, "Add command failed")
	},
}
