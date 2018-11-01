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

// unigramCmd represents the unigram command
var unigramCmd = &cobra.Command{
	Use:   "unigram",
	Short: "To rank Unigram Count",
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

		UnigramExecute(content, limit)
	},
}

//UnigramExecute execute main function
func UnigramExecute(filecontent string, n int) {
	docCollection := corpus.Ingest(filecontent)
	pmi := corpus.CreatePMI(&docCollection, getStopwords())
	pmi.TopCountList(n, "unigram", true)
}

func init() {
	rootCmd.AddCommand(unigramCmd)
}
