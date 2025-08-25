/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"todo_app/internal/todo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new tasks that have other properties like if completed or not. the date of the creation ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		existingTasks, err := Storage.Load()
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				existingTasks = &todo.Tasks{}
			} else {
				fmt.Printf("failed to load the file: %v", err)
				return
			}
		}
		existingTasks.Add(args[0])

		err = Storage.Save(*existingTasks)
		if err != nil {
			fmt.Printf("failed to save the file: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
