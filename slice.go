package yamlpath

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Slice struct {
	value []Element
}

func (s *Slice) Append(pointer Pointer, value interface{}) error {
	if !pointer.IsTarget() {
		index, err := strconv.Atoi(pointer.Head())
		if err != nil {
			return fmt.Errorf("cannot convert index of pointer to integer: %w", err)
		}
		return s.value[index].Append(pointer, value)
	}

	el, err := resolveValueElement(value)
	if err != nil {
		return err
	}

	s.value = append(s.value, el)
	return nil
}

func (s *Slice) Get(pointer Pointer) (interface{}, error) {
	if !pointer.HasChild() {
		return s.value, nil
	}

	index, err := strconv.Atoi(pointer.Head())
	if err != nil {
		return nil, fmt.Errorf("cannot convert index of pointer to integer: %w", err)
	}

	if index >= len(s.value) {
		return nil, fmt.Errorf("index out of bounds: %d", index)
	}

	return s.value[index].Get(pointer.Child())
}

func (s *Slice) Update(pointer Pointer, value interface{}) error {
	index, err := strconv.Atoi(pointer.Head())
	if err != nil {
		return fmt.Errorf("cannot convert index of pointer to integer: %w", err)
	}

	if index >= len(s.value) {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	val, err := resolveValueElement(value)
	if err != nil {
		return fmt.Errorf("cannot resolve value element: %w", err)
	}

	s.value[index] = val

	return nil
}

func (s *Slice) UnmarshalYAML(node *yaml.Node) error {
	s.value = make([]Element, len(node.Content))

	for i, n := range node.Content {
		var val Mixed

		if err := n.Decode(&val); err != nil {
			return err
		}

		s.value[i] = &val
	}

	return nil
}

func (s *Slice) MarshalYAML() (interface{}, error) {
	node := yaml.Node{
		Kind: yaml.SequenceNode,
	}

	for _, value := range s.value {
		valueNode := &yaml.Node{}
		if err := valueNode.Encode(value); err != nil {
			return nil, err
		}

		node.Content = append(node.Content, valueNode)
	}

	return &node, nil
}
