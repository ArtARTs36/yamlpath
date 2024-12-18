package yamlpath_test

import (
	"github.com/artarts36/yamlpath"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		Title    string
		Content  string
		Pointer  string
		Expected interface{}
	}{
		{
			Title: "get string",
			Content: `
user:
  name: John
`,
			Pointer:  "user.name",
			Expected: "John",
		},
		{
			Title: "get int",
			Content: `
user:
  age: 99
`,
			Pointer:  "user.age",
			Expected: 99,
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			doc, err := yamlpath.Unmarshall([]byte(c.Content))
			require.NoError(t, err)

			res, err := doc.Get(yamlpath.NewPointer(c.Pointer))
			require.NoError(t, err)
			assert.Equal(t, c.Expected, res)
		})
	}
}
