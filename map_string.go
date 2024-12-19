package yamlpath

import (
	"fmt"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"gopkg.in/yaml.v3"
)

type StringMap struct {
	value *orderedmap.OrderedMap[string, Element]
}

func (m *StringMap) Append(pointer *Pointer, value interface{}) error {
	if pointer == nil || !pointer.HasChild() {
		return fmt.Errorf("cannot append value to a nil pointer")
	}

	val, ok := m.value.Get(pointer.Head())
	if !ok {
		return fmt.Errorf("no value for key %q", pointer.Head())
	}

	return val.Append(pointer.Child(), value)
}

func (m *StringMap) Get(pointer *Pointer) (Element, error) {
	if pointer == nil || !pointer.HasChild() {
		return m, nil
	}

	val, ok := m.value.Get(pointer.Head())
	if !ok {
		return nil, fmt.Errorf("no value for key %q at pointer %q", pointer.Head(), pointer.Head())
	}

	return val.Get(pointer.Child())
}

func (m *StringMap) Update(pointer *Pointer, value interface{}) error {
	item, exists := m.value.Get(pointer.Head())
	if !exists {
		val, err := resolveValueElement(value)
		if err != nil {
			return err
		}
		m.value.Set(pointer.Head(), val)
		return nil
	}

	return item.Update(pointer.Child(), value)
}

func (m *StringMap) UnmarshalYAML(node *yaml.Node) error {
	m.value = orderedmap.New[string, Element]()

	for index := 0; index < len(node.Content); index += 2 {
		var key string
		var val Mixed

		if err := node.Content[index].Decode(&key); err != nil {
			return err
		}
		if err := node.Content[index+1].Decode(&val); err != nil {
			return err
		}

		m.value.Set(key, &val)
	}

	return nil
}

func (m *StringMap) MarshalYAML() (interface{}, error) {
	node := yaml.Node{
		Kind: yaml.MappingNode,
	}

	for pair := m.value.Oldest(); pair != nil; pair = pair.Next() {
		key, value := pair.Key, pair.Value

		keyNode := &yaml.Node{}

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
