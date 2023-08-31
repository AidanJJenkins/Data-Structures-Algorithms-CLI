package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func makeDir(cmd string) {
	today := time.Now().Format("01-02-06")
	newDirPath := filepath.Join(".", fmt.Sprintf("%s-%s", cmd, today))

	d := make(map[string]string)
	d["g1"] = "./template/g1"
	d["g2"] = "./template/g2"
	d["g3"] = "./template/g3"
	d["g4"] = "./template/g4"

	err := os.Mkdir(newDirPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %s\n", err)
		return
	}
	groupPath := d[cmd]
	err = copyFiles(groupPath, newDirPath)
	if err != nil {
		fmt.Printf("Error copying files: %s\n", err)
		return
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
