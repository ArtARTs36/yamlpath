package yamlpath

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"strconv"
)

type Str struct {
	value string
}

func (s *Str) Append(pointer *Pointer, value interface{}) error {
	if !pointer.IsTarget() {
		return errors.New("pointer is not target")
	}

	switch v := value.(type) {
	case string:
		s.value += v
	default:
		return errors.New("value is not string")
	}

	return nil
}

func (s *Str) Get(_ *Pointer) (Element, error) {
	return s, nil
}

func (s *Str) Update(pointer *Pointer, value interface{}) error {
	if pointer == nil || !pointer.IsTarget() {
		return errors.New("pointer is not target")
	}

	switch v := value.(type) {
	case string:
		s.value = v
	case int:
		s.value = strconv.Itoa(v)
	default:
		return fmt.Errorf("value is not string or number, got %T", value)
	}

	return nil
}

func (s *Str) UnmarshalYAML(node *yaml.Node) error {
	s.value = node.Value
	return nil
}

func (s *Str) MarshalYAML() (interface{}, error) {
	return s.value, nil
}

func (s *Str) Marshal() ([]byte, error) {
	return yaml.Marshal(s.value)
}

func (s *Str) AsScalar() (interface{}, error) {
	return s.value, nil
}
