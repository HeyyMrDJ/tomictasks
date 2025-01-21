package cmd

import (
	"fmt"
	"strconv"

	"github.com/heyymrdj/tomictasks/pkg/database"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "tag stuff",
}

var addTagCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tag to a node",
	Run: func(cmd *cobra.Command, args []string) {
		nodeID, _ := strconv.Atoi(args[0])
		tagID, _ := strconv.Atoi(args[1])
		database.AddTag(db, nodeID, tagID)
		fmt.Println("Tag added")
	},
}

var removeTagCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove tag to a node",
	Run: func(cmd *cobra.Command, args []string) {
		nodeID, _ := strconv.Atoi(args[0])
		tagID, _ := strconv.Atoi(args[1])
		database.RemoveTag(db, nodeID, tagID)
		fmt.Println("Tag remove")
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.AddCommand(addTagCmd)
	tagCmd.AddCommand(removeTagCmd)
}
