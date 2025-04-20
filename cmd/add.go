package cmd

import (
	"errors"
	"fmt"
	"passwordManger/internal/model"
	"passwordManger/internal/storage"
	"passwordManger/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [service] [password]",
	Short: "Add a new password entry",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			utils.CheckErr(errors.New("not enough arguments"), "Usage: add <service> <password> [note]")
		}

		service := args[0]
		password := args[1]

		entry := model.Entry{
			ID:        uuid.New().String(),
			Service:   service,
			Password:  password,
			Note:      fmt.Sprintf("Password for %s", service),
			CreatedAt: time.Now().Format(time.RFC3339),
		}
		err := storage.SaveEntry(entry)
		utils.CheckErr(err, "Failed to save entry")
		fmt.Println("✅ Saved password for", service)
	},
}
