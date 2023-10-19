package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func makeDir(cmd string) {
	today := time.Now().Format("01-02-06")
	newDirPath := filepath.Join(".", fmt.Sprintf("group%s_%s", cmd, today))

	d := make(map[string]string)
	d["1"] = "./template/easy/g1"
	d["2"] = "./template/easy/g2"
	d["3"] = "./template/easy/g3"
	d["4"] = "./template/easy/g4"

	easyCheck := fmt.Sprintf("./template/easy/g%s", cmd)

	if _, err := os.Stat(easyCheck); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", cmd)
	} else {
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

		fmt.Printf("practice problems for group: %s generated in: ./%s\n", cmd, newDirPath)
	}
}

func makeDirHard(cmd string) {
	today := time.Now().Format("01-02-06")
	newDirPath := filepath.Join(".", fmt.Sprintf("hard-group%s_%s", cmd, today))

	d := make(map[string]string)
	d["1"] = "./template/hard/g1h"
	d["2"] = "./template/hard/g2h"
	d["3"] = "./template/hard/g3h"
	d["4"] = "./template/hard/g4h"

	hardCheck := fmt.Sprintf("./template/hard/g%sh", cmd)

	if _, err := os.Stat(hardCheck); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", cmd)
	} else {
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

		fmt.Printf("hard practice problems for group: %s generated in: ./%s\n", cmd, newDirPath)
	}
}
