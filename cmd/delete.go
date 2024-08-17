package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/heyymrdj/tomictasks/pkg/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete resources",
}

var deleteListCmd = &cobra.Command{
	Use:   "list [listID]",
	Short: "deletes an existing list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Deleting list will also delete all tasks")
		fmt.Println("Do you want to proceed (y/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "y" || input == "yes" {
			fmt.Println("Deleting list...")
			name := args[0]
			id := database.GetListIDByName(db, name)
			database.DeleteList(db, id)
			fmt.Println("List deleted successfully")
		} else {
			fmt.Println("Operation canceled.")
			return
		}

	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "task [listID]",
	Short: "deletes an existing Task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if strings.Contains(args[0], "/") {
			task := strings.Split(args[0], "/")
			listName := task[0]
			taskName := task[1]
			listID := database.GetListIDByName(db, listName)
			taskID := database.GetTaskIDByName(db, taskName, listID)
			database.DeleteTask(db, taskID)
		} else {
			fmt.Println("Deleting task...")
			name := args[0]
			taskID := database.GetTaskIDByName(db, name, 1)
			database.DeleteTask(db, taskID)
			fmt.Println("Task deleted successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteTaskCmd)
	deleteCmd.AddCommand(deleteListCmd)
}
