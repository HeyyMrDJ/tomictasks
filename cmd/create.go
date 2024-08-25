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

var createNodeCmd = &cobra.Command{
	Use:   "node [node name]",
	Short: "Creates a node",
	Run: func(cmd *cobra.Command, args []string) {
		nodes := strings.Split(args[0], "/")
		if len(nodes) > 1 {
			var ids int
			for i, node := range nodes[:len(nodes)-1] {
				fmt.Println(node)
				n := database.GetNodeIDByName(db, nodes[i])
				ids = n[0].ID
			}
			parentID := ids
			path := "/" + strings.Join(nodes[:len(nodes)-1], "/") + "/"
			database.AddChildNode(db, parentID, nodes[len(nodes)-1], path)
		} else {
			database.CreateNode(db, args[0])
		}
	},
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
	createCmd.AddCommand(createNodeCmd)
}
