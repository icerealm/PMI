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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pmi/corpus"

	"github.com/spf13/cobra"
)

var stopwordFlg bool

var rootCmd = &cobra.Command{
	Use:   "pmi",
	Short: "CACM mutual word count",
	Long: `This CLI helps user to observe/summarize the various statistics in CACM corpus 
such as unigram wordcount, bigram wordcount`,
	Run: func(cmd *cobra.Command, args []string) {
		content, err := readFileContent(args)
		if err != nil {
			log.Fatalln(err)
		}
		PMIExecute(content)
	},
}

func init() {
	rootCmd.Flags().BoolVar(&stopwordFlg, "s", false, "enable stopword remove")
}

//Execute execute command here.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readFileContent(args []string) (string, error) {
	path := "cacm/cacm.all"
	// if len(args) < 1 {
	// 	return "", fmt.Errorf("no configuration file specified, using default path: %s", path)
	// }
	// path = args[0]
	// if path == "d" {
	// 	path = "cacm/cacm.all"
	// }
	content, err := readFile(path)
	if err != nil {
		return "", fmt.Errorf("no configuration file found with path specified: %s", path)
	}
	return content, nil
}

func getStopwords() []string {
	var stopwords []string
	if stopwordFlg {
		stopwordContent, err := readFile("cacm/common_words")
		if err != nil {
			log.Fatalln("error during get stopwords from file")
		}
		stopwords = corpus.GetStopWords(stopwordContent)
	}
	return stopwords
}

//PMIExecute execute main function
func PMIExecute(filecontent string) {
	docCollection := corpus.Ingest(filecontent)
	pmi := corpus.CreatePMI(&docCollection, getStopwords())
	pmi.TopCountList(10, "mutual", false)
}

func readFile(pathToFile string) (string, error) {
	dat, err := ioutil.ReadFile(pathToFile)
	return string(dat), err
}
