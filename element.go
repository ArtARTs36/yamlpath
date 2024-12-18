package yamlpath

type Element interface {
	Append(pointer Pointer, value interface{}) error
	Get(pointer Pointer) (interface{}, error)
	Update(pointer Pointer, value interface{}) error
	MarshalYAML() (interface{}, error)
}
