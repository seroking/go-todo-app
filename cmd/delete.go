/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task",
	Long:  `delete a task definitively from the file storage (json file for now)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

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

		err = actualTasks.Delete(int_arg)
		if err != nil {
			fmt.Printf("failed to delete the file: %v", err)
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
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
