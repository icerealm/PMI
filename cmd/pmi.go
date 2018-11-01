// Copyright © 2018 Jakkrit Sittiwerapong <EMAIL ADDRESS>
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
	"log"
	"pmi/corpus"
	"strconv"

	"github.com/spf13/cobra"
)

// pmiCmd represents the pmi command
var pmiCmd = &cobra.Command{
	Use:   "pmi",
	Short: "To observe pmi statistics",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := readFileContent(args)
		if err != nil {
			log.Fatalln(err)
		}
		n := "10"
		if len(args) > 0 {
			n = args[0]
		}
		limit, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln(err)
		}

		WordPMIExecute(content, limit)
	},
}

//WordPMIExecute execute main function
func WordPMIExecute(filecontent string, n int) {
	docCollection := corpus.Ingest(filecontent)
	pmi := corpus.CreatePMI(&docCollection, getStopwords())
	pmi.TopPMIList(n)
}
func init() {
	rootCmd.AddCommand(pmiCmd)
}
