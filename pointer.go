package yamlpath

import "strings"

type Pointer struct {
	head  int
	parts []string
}

func NewPointer(ptr string) *Pointer {
	parts := strings.Split(ptr, ".")
	return &Pointer{head: 0, parts: parts}
}

func (p *Pointer) Head() string {
	return p.parts[p.head]
}

func (p *Pointer) Child() *Pointer {
	return &Pointer{head: p.head + 1, parts: p.parts}
}

func (p *Pointer) HasChild() bool {
	return p.head < len(p.parts)
}

func (p *Pointer) IsTarget() bool {
	return len(p.parts) <= p.head
}
