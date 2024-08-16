package cmd

import (
	"fmt"
	"strings"

	"github.com/heyymrdj/tomictasks/pkg/database"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources",
}

var createTaskCmd = &cobra.Command{
	Use:   "task [listName] [title] [description] [dueDate]",
	Short: "Creates a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if strings.Contains(args[0], "/") {
			task := strings.Split(args[0], "/")
			listName := task[0]
			listID := database.GetListIDByName(db, listName)

			fmt.Println("Creating task on listName: ", listName)
			database.CreateTask(db, task[1], listID, "")
		} else {
			fmt.Println("Creating task on default list")
			database.CreateTask(db, args[0], 1, "")
			fmt.Println("Task created successfully")
		}
	},
}

var createListCmd = &cobra.Command{
	Use:   "list [listName]",
	Short: "Creates a new list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Creating list...", args[0])
		database.CreateList(db, args[0])
		fmt.Println("List created successfully")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createTaskCmd)
	createCmd.AddCommand(createListCmd)
}
