package cmd

import (
	"fmt"

	"github.com/heyymrdj/tomictasks/pkg/database"
	"github.com/spf13/cobra"
)

var defaultList = "default"

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get tasks and lists",
}

var getNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Get node",
	Run: func(cmd *cobra.Command, args []string) {
		nodeID, err := database.NodeWalk(db, args[0])
		if err != nil {
			fmt.Println(err)
			return
		} else {
			node := database.GetNode(db, nodeID)
			path := node.Path + node.Name + "/"
			nodes := database.GetChildren(db, path)
			fmt.Printf("%-30s\n", "NAME")
			if len(nodes) > 0 {
				fmt.Printf("%-30s\n\n", node.Name+"+")
				fmt.Printf("CHILDREN\n")
				for _, node := range nodes {
					path := node.Path + node.Name + "/"
					nodes := database.GetChildren(db, path)
					if len(nodes) > 0 {
						fmt.Printf("%-40s\n", node.Name+"+")
					} else {
						fmt.Printf("%-40s\n", node.Name)
					}
				}
			} else {
				fmt.Printf("%-30s\n\n", node.Name)
			}
		}
	},
}

var getNodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Get all root nodes",
	Run: func(cmd *cobra.Command, args []string) {
		nodes := database.GetNodes(db)
		fmt.Printf("%-45s\n", "ROOT NODES")
		for _, node := range nodes {
			path := "/" + node.Path + node.Name + "/"
			nodes := database.GetChildren(db, path)
			if len(nodes) > 0 {
				fmt.Printf("%-40s\n", node.Name+"+")
			} else {
				fmt.Printf("%-40s\n", node.Name)
			}
		}
	},
}

var getTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Get all tags",
	Run: func(cmd *cobra.Command, args []string) {
		database.GetTags(db)
	},
}

var getTagCmd = &cobra.Command{
	Use:   "tag [tag Key] [tag Value]",
	Short: "Get tag info and all nodes with this tag",
	Run: func(cmd *cobra.Command, args []string) {
		//tagID, _ := strconv.Atoi(args[0])
		tagID := database.GetTagIDByName(db, args[0], args[1])
		database.GetTag(db, tagID)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getNodeCmd)
	getCmd.AddCommand(getNodesCmd)
	getCmd.AddCommand(getTagsCmd)
	getCmd.AddCommand(getTagCmd)
}
