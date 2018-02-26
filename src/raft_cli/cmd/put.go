// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"raft_cli/utils"
	"os"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("put called")
		if utils.CheckNil("key", key) || utils.CheckNil("value", value) {
			os.Exit(1)
		}
		fmt.Printf("key: %s, value: %s\n", key, value)
		body := req.SendRequest("PUT", key, []byte(value))
		if body != nil || string(body) != ""{
			fmt.Println(string(body))
		}
	},
}

func init() {
	RootCmd.AddCommand(putCmd)
	putCmd.Flags().StringVarP(&key, "key", "k", "", "Key.")
	putCmd.Flags().StringVarP(&value, "value", "v", "", "Value.")
}
