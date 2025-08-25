/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print out the tasks in terminal",
	Long:  `print out the task in terminal , in form of a table. this command has no argument `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print called")

		actualTasks, err := Storage.Load()
		if err != nil {
			fmt.Printf("failed to load the file: %v", err)
			return
		}

		actualTasks.Print()

	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
