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

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createNodeCmd)
}
