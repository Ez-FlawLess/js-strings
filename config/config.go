package config

type Config struct {
	StringsEnumName string        `json:"strings-enum-name"`
	StringsEnumPath string        `json:"strings-enum-path"`
	StringFiles     []StringFiles `json:"string-files"`
}

type StringFiles struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
