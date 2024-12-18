package yamlpath

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func Unmarshall(content []byte) (*Document, error) {
	var doc Document

	err := yaml.Unmarshal(content, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func Get(content []byte, pointer string) (Element, error) {
	doc, err := Unmarshall(content)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling yaml: %w", err)
	}

	return doc.Get(NewPointer(pointer))
}

func GetScalar(content []byte, pointer string) (interface{}, error) {
	elem, err := Get(content, pointer)
	if err != nil {
		return nil, err
	}
	return elem.AsScalar()
}

func Update(content []byte, pointer string, value interface{}) error {
	doc, err := Unmarshall(content)
	if err != nil {
		return fmt.Errorf("error unmarshalling yaml: %w", err)
	}

	return doc.Update(NewPointer(pointer), value)
}

func Append(content []byte, pointer string, value interface{}) error {
	doc, err := Unmarshall(content)
	if err != nil {
		return fmt.Errorf("error unmarshalling yaml: %w", err)
	}

	return doc.Append(NewPointer(pointer), value)
}
