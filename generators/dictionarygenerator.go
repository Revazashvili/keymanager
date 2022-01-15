package generators

import (
	"bytes"
	"log"
	"text/template"
)

const (
	dictInsert = `{{range .}} 
INSERT INTO CIB_DICTIONARY (KEY_NAME,LANGUAGE_CODE,KEY_VALUE) VALUES('{{.KeyName}}','{{.LanguageCode}}','{{.KeyValue}}');{{end}}`
)

type DictionaryGenerator struct {
}

func NewDictionaryGenerator() Generator {
	return &DictionaryGenerator{}
}

func (dg *DictionaryGenerator) Generate(data interface{}) (result string, err error) {
	tmpl, err := template.New("DictionaryInsert").Parse(dictInsert)
	if err != nil {
		log.Fatal(err)
		return
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, data)
	if err != nil {
		log.Fatal(err)
		return
	}
	result = b.String()
	return
}
