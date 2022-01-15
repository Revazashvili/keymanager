package generators

import (
	"bytes"
	"log"
	"text/template"
)

const (
	dictKeyInsertTemplate = `{{range . }} 
INSERT INTO CIB_DICTIONARY_KEYS (KEY_NAME,KEY_DESCRIPTION,TOOLTIP,KEY_TYPE) VALUES('{{.KeyName}}','{{.KeyDescription}}','{{.TooltipKey}}','{{.KeyType}}');{{end}}`
)

type DictionaryKeyGenerator struct {
}

func NewDictionaryKeyGenerator() Generator {
	return &DictionaryKeyGenerator{}
}

func (dkg *DictionaryKeyGenerator) Generate(data interface{}) (result string, err error) {
	tmpl, err := template.New("DictionaryKeyInsert").Parse(dictKeyInsertTemplate)
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
