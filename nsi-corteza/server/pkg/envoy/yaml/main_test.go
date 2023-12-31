package yaml

import (
    "fmt"
    "os"
)

func parseDocument(name string) (*Document, error) {
	doc := &Document{}
	f, err := os.Open(fmt.Sprintf("testdata/%s.yaml", name))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return doc, yaml.NewDecoder(f).Decode(doc)
}
