/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a task",
	Long: `edit a task: you can only change the title of the task,
	 you choose the task by index in the first arg and the new title in the second arg`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")

		int_arg, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("failed to cast to integer type : %v", err)
			return
		}

		actualTasks, err := Storage.Load()
		if err != nil {
			fmt.Printf("failed to load the file: %v", err)
			return
		}

		err = actualTasks.Edit(int_arg, args[1])
		if err != nil {
			fmt.Printf("failed to edit the file: %v", err)
			return
		}

		err = Storage.Save(*actualTasks)
		if err != nil {
			fmt.Printf("failed to save the file: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
