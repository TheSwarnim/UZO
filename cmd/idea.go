/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"uzo/util"

	"github.com/spf13/cobra"
)

// ideaCmd represents the code command
var ideaCmd = &cobra.Command{
	Use:   "idea <zip_file_name>",
	Short: "It will open the directory in Intellij Idea",
	Long: `This command will help to open the unzipped folder
to Intellij Idea. In order for this command to work, Intellij 
idea should be installed in your system`,
	// Args: cobra.ExactArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) != 1 {
			// return errors.New("Please provide the file name to unzip and open in IDE")
			return errors.New("accept(s) 1 argument")
		}
		return nil
	},
	Example: `uzo code demo.zip`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string

		var argument string
		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}

		fileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err.Error())
		}

		if fileExists {
			fileName, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("File %v does not exist\n", argument)
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		util.Unzip(fileName, wd)
		os.Chdir(util.FilenameWithoutExtension(fileName))

		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		if err != nil {
			log.Fatal("Intellij Idea executable file not found in %PATH%")
			// fmt.Println("VS Code executable file not found in %PATH%")
		}
	},
}

func init() {
	rootCmd.AddCommand(ideaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ideaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ideaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add --file/-f flag to ideaCmd
	ideaCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "The file name to unzip and open in IDE")
}
