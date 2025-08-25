/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle a task",
	Long:  `Toggle a task : if completed assign it as not complete and vice-versa`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("toggle called")
		int_arg, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("failed to convert to integer : %v", err)
			return
		}
		actualTasks, err := Storage.Load()
		if err != nil {
			fmt.Printf("failed to load the file: %v", err)
			return
		}

		err = actualTasks.Toggle(int_arg)
		if err != nil {
			fmt.Printf("failed to toggle the task: %v", err)
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
	rootCmd.AddCommand(toggleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
