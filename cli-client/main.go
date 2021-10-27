/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gitlab.com/Prosp3r/ori/cli-client/cmd"
)

//"oriclient/cmd"

func main() {
	//cmd.Execute()
	for {
		fmt.Println("Enter a new command or type help : ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println("Enter a new command : ")
			//fmt.Println(scanner.Text())
			args := setArgs(scanner.Text())
			//oldArgs := []string{"cmd", args}
			//defer func() { os.Args = oldArgs }()
			oldArgs := make([]string, 0)
			oldArgs = append(oldArgs, "cmd")

			for _, entry := range args {
				oldArgs = append(oldArgs, entry)
			}
			os.Args = oldArgs
			cmd.Execute()
		}
	}
}

func setArgs(entry string) []string {
	arg := strings.Split(entry, " ")
	return arg
}
