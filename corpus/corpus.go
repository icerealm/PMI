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
	"regexp"
	"strings"
)

func contains(line string) bool {
	specialTokens := []string{".A", ".K", ".I", ".N", ".W"}
	for _, ele := range specialTokens {
		if ele == line {
			return true
		}
	}
	return false
}

//GetStopWords get stopwords
func GetStopWords(data string) []string {
	return strings.Split(data, "\n")
}

//Ingest data to collection.
func Ingest(data string) DocumentCollection {
	re := regexp.MustCompile("\\.I ")
	pDocuments := re.Split(data, -1)

	var docs DocumentCollection
	for _, pDocument := range pDocuments {
		lines := strings.Split(pDocument, "\n")
		var document []string
		for _, line := range lines {
			if !contains(line) && len(line) > 0 {
				document = append(document, line)
			}
		}
		if len(document) > 0 {
			docs.addElement(strings.Split(strings.Join(document, "\n"), ".X")[0])
		}
	}
	return docs
}

//Document document representation
type Document struct {
	documentID string
	content    string
	oneLine    string
	raw        string
}

//GetDocumentID return documentID
func (doc *Document) GetDocumentID() string {
	return doc.documentID
}

//GetContent return content
func (doc *Document) GetContent() string {
	return doc.content
}

//GetOneLine return oneLine
func (doc *Document) GetOneLine() string {
	return doc.oneLine
}

//DocumentCollection keep document collection
type DocumentCollection struct {
	documentLine []Document
}

func (docCollection *DocumentCollection) addElement(line string) {
	msg := strings.Split(line, ".T")
	id := strings.Replace(msg[0], "\n", " ", -1)
	content := strings.Split(msg[1], ".B")[0]
	oneline := strings.Replace(content, "\n", " ", -1)

	(*docCollection).documentLine = append(docCollection.documentLine, Document{
		documentID: strings.TrimSpace(id),
		content:    content,
		oneLine:    strings.Join(strings.Fields(oneline), " "),
		raw:        line,
	})
}

//Size count document size
func (docCollection *DocumentCollection) Size() int {
	return len(docCollection.documentLine)
}

//GetDocumentByIndex get document by id specified
func (docCollection *DocumentCollection) GetDocumentByIndex(idx int) Document {
	return docCollection.documentLine[idx]
}

func (docCollection *DocumentCollection) getDocumentByField(idx int, field string) string {
	if field == "raw" {
		return docCollection.documentLine[idx].raw
	}
	if field == "content" {
		return docCollection.documentLine[idx].content
	}
	if field == "oneline" {
		return docCollection.documentLine[idx].oneLine
	}
	return docCollection.documentLine[idx].documentID
}

//Print show all documents in collection.
func (docCollection DocumentCollection) Print() {
	fmt.Println(docCollection.documentLine)
}

//Top show head n documents.
func (docCollection DocumentCollection) Top(nLines int, field string) {
	for i := 0; i < nLines; i++ {
		fmt.Println(docCollection.getDocumentByField(i, field))
	}
}

//Tail show tail n documents
func (docCollection DocumentCollection) Tail(nLines int, field string) {
	size := len(docCollection.documentLine)
	for i := size - nLines; i < size; i++ {
		fmt.Println(docCollection.getDocumentByField(i, field))
	}
}
