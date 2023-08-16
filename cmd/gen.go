/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func copyTemplateFiles(templatePath, newPath string) error {
	return filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(templatePath, path)
		newPath := filepath.Join(newPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(newPath, info.Mode())
		}

		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(newPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		return err
	})
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new day directory for practice problems",
	Run: func(cmd *cobra.Command, args []string) {
		templateDir := "./template"

		dayName := time.Now().Format("01-02-06")
		newDirPath := filepath.Join(".", "day-"+dayName)

		err := os.Mkdir(newDirPath, 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %s\n", err)
			return
		}

		err = copyTemplateFiles(templateDir, newDirPath)
		if err != nil {
			fmt.Printf("Error copying files: %s\n", err)
			return
		}

		fmt.Printf("Day %s practice problems generated in: %s\n", dayName, newDirPath)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
