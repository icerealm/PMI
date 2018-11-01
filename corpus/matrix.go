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
	"sort"
	"strings"
)

// WordDocument map data structure beween word and documentIds
type WordDocument struct {
	WordMatrix map[string]*Posting
	Corpus     *DocumentCollection
}

//Posting posting
type Posting struct {
	DocumentIdx []int
	Word        []string
}

func indexContains(indexes []int, idx int) bool {
	for _, val := range indexes {
		if val == idx {
			return true
		}
	}
	return false
}

func isExistingWord(qword string, words []string) bool {
	for _, word := range words {
		if word == qword {
			return true
		}
	}
	return false
}

func (posting *Posting) append(idx int) {
	if !indexContains(posting.DocumentIdx, idx) {
		posting.DocumentIdx = append(posting.DocumentIdx, idx)
	}
}

func (posting *Posting) appendWithIdxAndWord(idx int, word string) {
	if !indexContains(posting.DocumentIdx, idx) {
		posting.DocumentIdx = append(posting.DocumentIdx, idx)
		posting.Word = append(posting.Word, word)
	}
}

func filterWords(words []string, stopwords []string) []string {
	var cleanWords []string
	if len(stopwords) > 0 {
		for _, pword := range words {
			word := strings.TrimSuffix(pword, ".")
			if !isExistingWord(word, stopwords) {
				cleanWords = append(cleanWords, word)
			}
		}
	} else {
		for _, pword := range words {
			cleanWords = append(cleanWords, strings.TrimSuffix(pword, "."))
		}
	}

	return cleanWords
}

func appendUnigramMatrix(wordMat map[string]*Posting, doc Document, docIndex int, stopwords []string) {
	words := filterWords(strings.Split(doc.GetOneLine(), " "), stopwords)
	for _, word := range words {
		key := strings.ToLower(word)
		if wordMat[key] == nil {
			posting := Posting{}
			posting.append(docIndex)
			wordMat[key] = &posting
		} else if wordMat[key] != nil {
			wordMat[key].append(docIndex)
		}
	}
}

func appendBigramMatrix(wordMat map[string]*Posting, doc Document, docIndex int, stopwords []string) {
	words := filterWords(strings.Split(doc.GetOneLine(), " "), stopwords)
	for idx := 1; idx < len(words); idx++ {
		bigram := []string{strings.ToLower(words[idx-1]), strings.ToLower(words[idx])}
		key := strings.Join(bigram, " ")
		if wordMat[key] == nil {
			posting := Posting{}
			posting.appendWithIdxAndWord(docIndex, key)
			wordMat[key] = &posting
		} else if wordMat[key] != nil {
			wordMat[key].appendWithIdxAndWord(docIndex, key)
		}
	}
}

func appendMutualMatrix(wordMat map[string]*Posting, doc Document, docIndex int, stopwords []string) {
	words := filterWords(strings.Split(doc.GetOneLine(), " "), stopwords)
	for idx := 1; idx < len(words); idx++ {
		bigram := []string{strings.ToLower(words[idx-1]), strings.ToLower(words[idx])}
		sort.Strings(bigram) //nutural bi-words
		key := strings.Join(bigram, " ")

		if wordMat[key] == nil {
			posting := Posting{}
			posting.appendWithIdxAndWord(docIndex, key)
			wordMat[key] = &posting
		} else if wordMat[key] != nil {
			wordMat[key].appendWithIdxAndWord(docIndex, key)
		}
	}
}

//CreateUnigramWordMatrix return unigram matrix
func CreateUnigramWordMatrix(corpus *DocumentCollection, stopwords []string) WordDocument {
	matrix := make(map[string]*Posting)
	size := corpus.Size()
	for i := 0; i < size; i++ {
		doc := corpus.GetDocumentByIndex(i)
		appendUnigramMatrix(matrix, doc, i, stopwords)
	}
	return WordDocument{
		WordMatrix: matrix,
		Corpus:     corpus,
	}
}

//CreateBigramMatrix return bigram matrix
func CreateBigramMatrix(corpus *DocumentCollection, stopwords []string) WordDocument {
	matrix := make(map[string]*Posting)
	size := corpus.Size()
	for i := 0; i < size; i++ {
		doc := corpus.GetDocumentByIndex(i)
		appendBigramMatrix(matrix, doc, i, stopwords)
	}
	return WordDocument{
		WordMatrix: matrix,
		Corpus:     corpus,
	}
}

//CreateMutualWordMatrix return mutual word matrix
func CreateMutualWordMatrix(corpus *DocumentCollection, stopwords []string) WordDocument {
	matrix := make(map[string]*Posting)
	size := corpus.Size()
	for i := 0; i < size; i++ {
		doc := corpus.GetDocumentByIndex(i)
		appendMutualMatrix(matrix, doc, i, stopwords)
	}
	return WordDocument{
		WordMatrix: matrix,
		Corpus:     corpus,
	}
}
