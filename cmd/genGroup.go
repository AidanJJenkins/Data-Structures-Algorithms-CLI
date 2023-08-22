package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func makeDir(cmd string) {
	var groupDirectories = map[string][]string{
		"g1": {"./template/linearSearchList", "./template/binarySearchList", "./template/twoCrystalBalls", "./template/singlyLinkedList", "./template/stack", "./template/queue"},
	}
	if groupDirectories[cmd] == nil {
		fmt.Println("group does not exist, possible groups are: g1, g2, g3, and g4")
	}

	today := time.Now().Format("01-02-06")
	newDirPath := filepath.Join(".", fmt.Sprintf("%s-%s", cmd, today))

	err := os.Mkdir(newDirPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %s\n", err)
		return
	}

	for _, dir := range groupDirectories[cmd] {
		err = copyFiles(dir, newDirPath)
		if err != nil {
			fmt.Printf("Error copying files: %s\n", err)
			return
		}
	}

	fmt.Printf("Day %s practice problems generated in: %s\n", today, newDirPath)
}

var group string
var genGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "generate one group of the problems",
	Run: func(cmd *cobra.Command, args []string) {
		makeDir(group)
	},
}

func init() {
	rootCmd.AddCommand(genGroupCmd)
	genGroupCmd.Flags().StringVarP(&group, "group", "g", "", "any group of problems to be generated")
	genGroupCmd.MarkFlagRequired("group")
}
