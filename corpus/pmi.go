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

package corpus

import (
	"fmt"
	"sort"
	"strings"
)

//PMI represent pmi
type PMI struct {
	unigramMatrix *WordDocument
	bigramMatrix  *WordDocument
	mutualMatrix  *WordDocument
}

//CreatePMI create pmi instance
func CreatePMI(corpus *DocumentCollection, stopwords []string) *PMI {
	unigramMat := CreateUnigramWordMatrix(corpus, stopwords)
	bigramMat := CreateBigramMatrix(corpus, stopwords)
	mutualMat := CreateMutualWordMatrix(corpus, stopwords)

	return &PMI{
		unigramMatrix: &unigramMat,
		bigramMatrix:  &bigramMat,
		mutualMatrix:  &mutualMat,
	}
}

//TopCountList to show top rank
func (pmi PMI) TopCountList(n int, option string, needStatistics bool) {
	var wMat map[string]*Posting
	var totalDocuments int
	if option == "unigram" {
		wMat = pmi.unigramMatrix.WordMatrix
		totalDocuments = pmi.unigramMatrix.Corpus.Size()
	}
	if option == "bigram" {
		wMat = pmi.bigramMatrix.WordMatrix
		totalDocuments = pmi.bigramMatrix.Corpus.Size()
	}
	if option == "mutual" {
		wMat = pmi.mutualMatrix.WordMatrix
		totalDocuments = pmi.mutualMatrix.Corpus.Size()
	}

	type pair struct {
		word  string
		count int
		prob  float64
	}
	var pairList []pair
	for k := range wMat {
		count := len(wMat[k].DocumentIdx)
		wprob := float64(count) / float64(totalDocuments)
		pairList = append(pairList, pair{
			word:  k,
			count: count,
			prob:  wprob,
		})
	}
	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].count > pairList[j].count
	})

	if needStatistics {
		for idx := 0; idx < n; idx++ {
			fmt.Println(fmt.Sprintf("'%v' : %v, prob=%v", pairList[idx].word, pairList[idx].count, pairList[idx].prob))
		}
	} else {
		for idx := 0; idx < n; idx++ {
			fmt.Println(fmt.Sprintf("'%v' : %v", pairList[idx].word, pairList[idx].count))
		}
	}

}

//TopPMIList observe pmi
func (pmi PMI) TopPMIList(n int) {
	uniMat := pmi.unigramMatrix.WordMatrix
	mutualMat := pmi.mutualMatrix.WordMatrix

	totalDocuments := float64(pmi.unigramMatrix.Corpus.Size())
	type wordpmi struct {
		word string
		pmi  float64
	}
	var wordpmiList []wordpmi

	for k := range mutualMat {
		mutualProp := float64(len(mutualMat[k].DocumentIdx)) / totalDocuments
		kwords := strings.Split(k, " ")
		firstwordProp := float64(len(uniMat[kwords[0]].DocumentIdx)) / totalDocuments
		secondwordProp := float64(len(uniMat[kwords[1]].DocumentIdx)) / totalDocuments
		wpmi := mutualProp / (firstwordProp * secondwordProp)
		wordpmiList = append(wordpmiList, wordpmi{
			word: k,
			pmi:  wpmi,
		})
	}

	// sort.Slice(wordpmiList, func(i, j int) bool {
	// 	diff := wordpmiList[i].pmi - wordpmiList[j].pmi
	// 	return diff > 0.0000000001
	// })

	for idx := 0; idx < n; idx++ {
		fmt.Println(wordpmiList[idx].word, ":", wordpmiList[idx].pmi)
	}
}

//GetUnigramMatrix get unigram matrix
func (pmi PMI) GetUnigramMatrix() WordDocument {
	return *pmi.unigramMatrix
}

//GetBigramMatrix get bigram matrix
func (pmi PMI) GetBigramMatrix() WordDocument {
	return *pmi.bigramMatrix
}

//GetMutualMatrix get mutual matrix
func (pmi PMI) GetMutualMatrix() WordDocument {
	return *pmi.mutualMatrix
}
