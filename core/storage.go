package keymanager

type OracleStorage struct {
}

func NewOracleStorage() *OracleStorage {
	return &OracleStorage{}
}

func (os *OracleStorage) GetDictionaryKeys(p string) (r []DictionaryKey, err error) {
	r = []DictionaryKey{
		{"KeyName1", "KeyDescription1", "TooltipKey1", "KeyType1"},
		{"KeyName2", "KeyDescription2", "TooltipKey2", "KeyType2"},
		{"KeyName3", "KeyDescription3", "TooltipKey3", "KeyType3"},
		{"KeyName4", "KeyDescription4", "TooltipKey4", "KeyType4"},
	}
	return
}

func (os *OracleStorage) GetDictionaries(p string) (r []Dictionary, err error) {
	r = []Dictionary{
		{"KeyName1", "GE", "KeyValue1"},
		{"KeyName2", "EN", "KeyValue2"},
		{"KeyName3", "TR", "KeyValue3"},
		{"KeyName4", "RU", "KeyValue4"},
	}
	return
}
