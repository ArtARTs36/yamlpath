package yamlpath

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strconv"
)

type Scalar struct {
	value Element
}

func (s *Scalar) Append(pointer Pointer, value interface{}) error {
	return s.value.Append(pointer, value)
}

func (s *Scalar) Get(pointer Pointer) (Element, error) {
	return s.value.Get(pointer)
}

func (s *Scalar) Update(pointer Pointer, value interface{}) error {
	return s.value.Update(pointer, value)
}

func (s *Scalar) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.ScalarNode {
		return errors.New("expected scalar node")
	}

	if intVal, err := strconv.Atoi(node.Value); err == nil {
		s.value = &Int{intVal}
		return nil
	}

	s.value = &Str{node.Value}

	return nil
}

func (s *Scalar) MarshalYAML() (interface{}, error) {
	return s.value.MarshalYAML()
}

func (s *Scalar) Marshal() ([]byte, error) {
	return yaml.Marshal(s.value)
}

func (s *Scalar) AsScalar() (interface{}, error) {
	return s.value.AsScalar()
}
