package yamlpath

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Document struct {
	doc StringMap
}

func (d *Document) Append(pointer *Pointer, value interface{}) error {
	return d.doc.Append(pointer, value)
}

func (d *Document) Get(pointer *Pointer) (Element, error) {
	return d.doc.Get(pointer)
}

func (d *Document) Update(pointer *Pointer, value interface{}) error {
	if pointer.IsTarget() {
		return fmt.Errorf("pointer %q invalid: must be non-root element", pointer)
	}

	return d.doc.Update(pointer, value)
}

func (d *Document) UnmarshalYAML(node *yaml.Node) error {
	d.doc = StringMap{}
	return node.Decode(&d.doc)
}

func (d *Document) MarshalYAML() (interface{}, error) {
	return d.doc.MarshalYAML()
}

func (d *Document) Marshal() ([]byte, error) {
	return yaml.Marshal(d)
}

func (d *Document) AsScalar() (interface{}, error) {
	return nil, ErrElementNoScalar
}
