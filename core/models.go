package keymanager

type DictionaryKey struct {
	KeyName        string
	KeyDescription string
	TooltipKey     string
	KeyType        string
}

type Dictionary struct {
	KeyName      string
	LanguageCode string
	KeyValue     string
}
