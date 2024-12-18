package yamlpath

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Mixed struct {
	elem Element
}

func (m *Mixed) Append(pointer Pointer, value interface{}) error {
	return m.elem.Append(pointer, value)
}

func (m *Mixed) Get(pointer Pointer) (Element, error) {
	return m.elem.Get(pointer)
}

func (m *Mixed) Update(pointer Pointer, value interface{}) error {
	if pointer.IsTarget() {
		var err error
		m.elem, err = resolveValueElement(value)
		if err != nil {
			return err
		}
	}

	return m.elem.Update(pointer, value)
}

func (m *Mixed) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind { //nolint:exhaustive // not need
	case yaml.MappingNode:
		m.elem = &StringMap{}
	case yaml.ScalarNode:
		m.elem = &Scalar{}
	case yaml.SequenceNode:
		m.elem = &Slice{}
	default:
		return fmt.Errorf("node of type %v cannot be unmarshalled as mixed", node.Kind)
	}

	return node.Decode(m.elem)
}

func (m *Mixed) MarshalYAML() (interface{}, error) {
	return m.elem.MarshalYAML()
}

func (m *Mixed) Marshal() ([]byte, error) {
	return yaml.Marshal(m.elem)
}

func (m *Mixed) AsScalar() (interface{}, error) {
	return m.elem.AsScalar()
}
