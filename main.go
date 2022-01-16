/*
Copyright © 2022 levanrevazashvili@gmail.com

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
package main

import "github.com/Revazashvili/keymanager/cmd"

func main() {
	cmd.Execute()
}

//package main
//
//import (
//	"fmt"
//	keymanager "github.com/Revazashvili/keymanager/core"
//	"github.com/Revazashvili/keymanager/generators"
//	"io/fs"
//	"io/ioutil"
//	"os"
//)
//
//const (
//	path = "/home/revazashvili/go/src/github.com/Revazashvili/keymanager/script.sql"
//)
//
//var dictKeyInsertModel = []keymanager.DictionaryKey{
//	{"KeyName1", "KeyDescription1", "TooltipKey1", "KeyType1"},
//	{"KeyName2", "KeyDescription2", "TooltipKey2", "KeyType2"},
//	{"KeyName3", "KeyDescription3", "TooltipKey3", "KeyType3"},
//	{"KeyName4", "KeyDescription4", "TooltipKey4", "KeyType4"},
//}
//
//var dictInsertModel = []keymanager.Dictionary{
//	{"KeyName1", "GE", "KeyValue1"},
//	{"KeyName2", "EN", "KeyValue2"},
//	{"KeyName3", "TR", "KeyValue3"},
//	{"KeyName4", "RU", "KeyValue4"},
//}
//
//func main() {
//	dictKeyGen := generators.NewDictionaryKeyGenerator()
//	r, err := dictKeyGen.Generate(dictKeyInsertModel)
//	if err != nil {
//		panic(err)
//	}
//	dictGen := generators.NewDictionaryGenerator()
//	r2, err := dictGen.Generate(dictInsertModel)
//	if err != nil {
//		panic(err)
//	}
//	err = os.Chmod(path, 0777)
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = ioutil.WriteFile(path, []byte(r+"\n"+r2), fs.ModeAppend)
//	if err != nil {
//		panic(err)
//	}
//}
