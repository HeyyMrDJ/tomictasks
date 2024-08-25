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

var deleteNodeCmd = &cobra.Command{
	Use:   "node [nodeID]",
	Short: "Deletes node and children",
	Run: func(cmd *cobra.Command, args []string) {
		nodes := strings.Split(args[0], "/")
		path := strings.Join(nodes[:len(nodes)-1], "/") + "/"
		switch path {
		case "/":
		default:
			path = "/" + path
		}
		name := nodes[len(nodes)-1]
		nodeID := database.GetNodeIDByPath(db, name, path)

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Deleting list will also delete all tasks")
		fmt.Println("Do you want to proceed (y/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "y" || input == "yes" {
			fmt.Println("Deleting list...")
			database.DeleteNode(db, nodeID)
		} else {
			fmt.Println("Operation canceled.")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteNodeCmd)
}
