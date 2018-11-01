package test

import (
	"io/ioutil"
	"pmi/corpus"
	"testing"
)

var docCollection corpus.DocumentCollection
var stopwords []string

func init() {
	data, _ := ioutil.ReadFile("./cacm.all")
	docCollection = corpus.Ingest(string(data))
}

func TestIngest(t *testing.T) {
	if docCollection.Size() != 3204 {
		t.Fail()
	}
}

func TestGetDocumentByIndex(t *testing.T) {
	d1 := docCollection.GetDocumentByIndex(3203)
	if d1.GetDocumentID() != "3204" {
		t.Fail()
	}
	d2 := docCollection.GetDocumentByIndex(0)
	if d2.GetDocumentID() != "1" {
		t.Fail()
	}
}

func TestCreateUnigramWordMatrix(t *testing.T) {
	w := corpus.CreateUnigramWordMatrix(&docCollection, stopwords)
	if len(w.WordMatrix["queue"].DocumentIdx) == 0 {
		t.Fail()
	}
	// fmt.Println(w.WordMatrix["queue"].DocumentIdx)
}

func TestCreateBigramWordMatrix(t *testing.T) {
	w := corpus.CreateBigramMatrix(&docCollection, stopwords)
	if len(w.WordMatrix["a function"].DocumentIdx) == 0 {
		t.Fail()
	}
	// fmt.Println(w.WordMatrix["a function"].DocumentIdx)
}

func TestCreatePMI(t *testing.T) {
	pmi := corpus.CreatePMI(&docCollection, stopwords)
	if len(pmi.GetBigramMatrix().WordMatrix["a function"].DocumentIdx) == 0 {
		t.Fail()
	}
	if len(pmi.GetUnigramMatrix().WordMatrix["queue"].DocumentIdx) == 0 {
		t.Fail()
	}
	// pmi.TopPMIList(10)
}
