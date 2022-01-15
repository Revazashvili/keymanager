package main

import (
	"github.com/Revazashvili/keymanager/generators"
	"github.com/Revazashvili/keymanager/models"
	"io/fs"
	"io/ioutil"
)

const (
	path = "/home/revazashvili/go/src/github.com/Revazashvili/keymanager/script.sql"
)

var dictKeyInsertModel = []models.DictionaryKey{
	{"KeyName1", "KeyDescription1", "TooltipKey1", "KeyType1"},
	{"KeyName2", "KeyDescription2", "TooltipKey2", "KeyType2"},
	{"KeyName3", "KeyDescription3", "TooltipKey3", "KeyType3"},
	{"KeyName4", "KeyDescription4", "TooltipKey4", "KeyType4"},
}

var dictInsertModel = []models.Dictionary{
	{"KeyName1", "GE", "KeyValue1"},
	{"KeyName2", "EN", "KeyValue2"},
	{"KeyName3", "TR", "KeyValue3"},
	{"KeyName4", "RU", "KeyValue4"},
}

func main() {
	dictKeyGen := generators.NewDictionaryKeyGenerator()
	r, err := dictKeyGen.Generate(dictKeyInsertModel)
	if err != nil {
		panic(err)
	}
	dictGen := generators.NewDictionaryGenerator()
	r2, err := dictGen.Generate(dictInsertModel)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, []byte(r+"\n"+r2), fs.ModeAppend)
	if err != nil {
		panic(err)
	}
}
