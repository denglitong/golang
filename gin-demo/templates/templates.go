package templates

import (
	"html/template"
	"io/ioutil"
	"strings"
)

const (
	TPL_SUFFIX = ".htm"
)

func LoadTemplates() (*template.Template, error) {
	t := template.New("default")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, TPL_SUFFIX) {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
