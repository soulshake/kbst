/*
Copyright © 2020 Kubestack <hello@kubestack.com>

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
	"log"

	"github.com/kbst/kbst/cli"
	"github.com/spf13/cobra"
)

var release string
var gitRef string

// repositoryInitCmd represents the repository init command
var repositoryInitCmd = &cobra.Command{
	Use:   "init <starter>",
	Short: "Scaffold a new repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		starter := args[0]
		err := cli.RepoInit(starter, release, gitRef, path)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	repositoryCmd.AddCommand(repositoryInitCmd)

	repositoryInitCmd.Flags().StringVarP(&release, "release", "r", "latest", "desired release version")
	repositoryInitCmd.Flags().StringVar(&gitRef, "gitref", "", "git ref to download a dev artifact")
	repositoryInitCmd.Flags().MarkHidden("gitref")
}
