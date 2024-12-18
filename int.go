package yamlpath

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
)

type Int struct {
	value int
}

func (i *Int) Append(_ Pointer, value interface{}) error {
	intValue, ok := value.(int)
	if !ok {
		return errors.New("cannot append int value")
	}

	i.value += intValue

	return nil
}

func (i *Int) Get(_ Pointer) (interface{}, error) {
	return i.value, nil
}

func (i *Int) Update(pointer Pointer, value interface{}) error {
	if !pointer.IsTarget() {
		return errors.New("pointer is not target")
	}

	switch v := value.(type) {
	case int:
		i.value = v
	default:
		return fmt.Errorf("value of type %T is not number", value)
	}

	return nil
}

func (i *Int) UnmarshalYAML(node *yaml.Node) error {
	return node.Decode(&i.value)
}

func (i *Int) MarshalYAML() (interface{}, error) {
	return i.value, nil
}
