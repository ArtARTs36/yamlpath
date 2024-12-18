package yamlpath

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type StringMap struct {
	value map[string]Element
}

func (m *StringMap) Append(pointer Pointer, value interface{}) error {
	if !pointer.HasChild() {
		return fmt.Errorf("cannot append value to a nil pointer")
	}

	_, ok := m.value[pointer.Head()]
	if !ok {
		return fmt.Errorf("no value for key %q", pointer.Head())
	}

	return m.value[pointer.Head()].Append(pointer.Child(), value)
}

func (m *StringMap) Get(pointer Pointer) (Element, error) {
	if !pointer.HasChild() {
		return m, nil
	}

	val, ok := m.value[pointer.Head()]
	if !ok {
		return nil, fmt.Errorf("no value for key %q at pointer %q", pointer.Head(), pointer.Head())
	}

	return val.Get(pointer.Child())
}

func (m *StringMap) Update(pointer Pointer, value interface{}) error {
	if _, exists := m.value[pointer.Head()]; !exists {
		val, err := resolveValueElement(value)
		if err != nil {
			return err
		}
		m.value[pointer.Head()] = val
		return nil
	}

	return m.value[pointer.Head()].Update(pointer.Child(), value)
}

func (m *StringMap) UnmarshalYAML(node *yaml.Node) error {
	m.value = make(map[string]Element)

	for index := 0; index < len(node.Content); index += 2 {
		var key string
		var val Mixed

		if err := node.Content[index].Decode(&key); err != nil {
			return err
		}
		if err := node.Content[index+1].Decode(&val); err != nil {
			return err
		}

		m.value[key] = &val
	}

	return nil
}

func (m *StringMap) MarshalYAML() (interface{}, error) {
	node := yaml.Node{
		Kind: yaml.MappingNode,
	}

	for key, value := range m.value {
		keyNode := &yaml.Node{}

		// serialize key to yaml, then deserialize it back into the node
		// this is a hack to get the correct tag for the key
		if err := keyNode.Encode(key); err != nil {
			return nil, err
		}

		valueNode := &yaml.Node{}
		if err := valueNode.Encode(value); err != nil {
			return nil, err
		}

		node.Content = append(node.Content, keyNode, valueNode)
	}

	return &node, nil
}

func (m *StringMap) Marshal() ([]byte, error) {
	return yaml.Marshal(m.value)
}

func (m *StringMap) AsScalar() (interface{}, error) {
	return m.value, nil
}
