/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/heyymrdj/tomictasks/pkg/database"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update resources",
}

var updateListCmd = &cobra.Command{
	Use:   "list [listID]",
	Short: "updates an existing list",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Updating list...")
		id, _ := strconv.Atoi(args[0])
		database.UpdateList(db, id, args[1])
		fmt.Println("List updated successfully")
	},
}

var updateTaskCmd = &cobra.Command{
	Use:   "task [listID]",
	Short: "updates an existing task",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Updating task...")
		id, _ := strconv.Atoi(args[0])
		database.UpdateTask(db, id, args[1])
		fmt.Println("Task updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateListCmd)
	updateCmd.AddCommand(updateTaskCmd)
}
