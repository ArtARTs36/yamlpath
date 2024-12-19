package yamlpath

import (
	"errors"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Bool struct {
	value bool
}

func (s *Bool) Append(_ *Pointer, _ interface{}) error {
	return errors.New("cannot append bool value")
}

func (s *Bool) Get(_ *Pointer) (Element, error) {
	return s, nil
}

func (s *Bool) Update(pointer *Pointer, value interface{}) error {
	if !pointer.IsTarget() {
		return errors.New("pointer is not target")
	}

	switch v := value.(type) {
	case bool:
		s.value = v
	case string:
		if v == "true" || v == "TRUE" {
			s.value = true
		} else if v == "false" || v == "FALSE" {
			s.value = false
		}
		return fmt.Errorf("cannot update boolean value of string %q", v)
	default:
		return fmt.Errorf("cannot update boolean value of type %T", v)
	}

	return nil
}

func (s *Bool) UnmarshalYAML(node *yaml.Node) error {
	return node.Decode(&s.value)
}

func (s *Bool) MarshalYAML() (interface{}, error) {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: strconv.FormatBool(s.value),
	}, nil
}

func (s *Bool) Marshal() ([]byte, error) {
	return yaml.Marshal(s.value)
}

func (s *Bool) AsScalar() (interface{}, error) {
	return s.value, nil
}
