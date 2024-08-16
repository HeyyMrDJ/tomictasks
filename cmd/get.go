package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/heyymrdj/tomictasks/pkg/database"
	"github.com/spf13/cobra"
)

var defaultList = "default"

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get tasks and lists",
}

var getListsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Gets all lists",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement logic to get all lists
		lists := database.GetLists(db, "")
		fmt.Printf("%-45s %-10s\n", "NAME", "ID")
		for _, list := range lists {
			fmt.Printf("%-45s %-10d\n", list.Name, list.ID)
		}
	},
}

var getListCmd = &cobra.Command{
	Use:   "list [listName]",
	Short: "Gets all tasks for a specified list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Implement logic to get all tasks for a specified list
		fmt.Printf("%-45s %-10s\n", "NAME", "ID")
		listID := database.GetListIDByName(db, args[0])
		tasks := database.GetList(db, listID)
		for _, task := range tasks {
			fmt.Printf("%-45s %-10d\n", task.Title, task.ID)
		}
	},
}

var getAllTasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Gets all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement logic to get all tasks
		tasks := database.GetTasks(db, "")
		fmt.Printf("%-45s %-10s\n", "NAME", "LISTNAME")
		for _, task := range tasks {
			fmt.Printf("%-45s %-10s\n", task.Title, task.ListName)
		}

	},
}

var getDueTasksCmd = &cobra.Command{
	Use:   "due",
	Short: "Gets all tasks coming due or past due",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement logic to get due or past due tasks
		fmt.Println("Getting tasks due or past due...")
	},
}

var getTaskCmd = &cobra.Command{
	Use:   "task [listName/taskName] or task [taskName]",
	Short: "Creates a task on aspecified list, or from the default list if no list is specified",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if strings.Contains(args[0], "/") {
			task := strings.Split(args[0], "/")
			listID, _ := strconv.Atoi(task[0])
			taskID := database.GetTaskIDByName(db, task[1], listID)
			database.ReadTask(db, taskID)
		} else {
			taskID := database.GetTaskIDByName(db, args[0], 1)
			database.ReadTask(db, taskID)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getListsCmd)
	getCmd.AddCommand(getListCmd)
	getCmd.AddCommand(getAllTasksCmd)
	getCmd.AddCommand(getDueTasksCmd)
	getCmd.AddCommand(getTaskCmd)
}
