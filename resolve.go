package yamlpath

import "fmt"

func resolveValueElement(value interface{}) (Element, error) {
	switch v := value.(type) {
	case string:
		return &Str{v}, nil
	case int:
		return &Int{v}, nil
	case bool:
		return &Bool{v}, nil
	}
	return nil, fmt.Errorf("cannot resolve value of type %T: %v", value, value)
}
