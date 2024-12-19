package yamlpath

import "errors"

type Element interface {
	Append(pointer *Pointer, value interface{}) error
	Get(pointer *Pointer) (Element, error)
	Update(pointer *Pointer, value interface{}) error
	MarshalYAML() (interface{}, error)
	Marshal() ([]byte, error)
	// AsScalar can return ErrElementNoScalar
	AsScalar() (interface{}, error)
}

var ErrElementNoScalar = errors.New("element is not a scalar")
