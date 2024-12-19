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
		{
			Title: "get bool true",
			Content: `
user:
   active: true`,
			Pointer:  "user.active",
			Expected: true,
		},
		{
			Title: "get bool false",
			Content: `
user:
   active: false`,
			Pointer:  "user.active",
			Expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			doc, err := yamlpath.Unmarshal([]byte(c.Content))
			require.NoError(t, err)

			elem, err := doc.Get(yamlpath.NewPointer(c.Pointer))
			require.NoError(t, err)

			res, err := elem.AsScalar()
			assert.Equal(t, c.Expected, res)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Run("update map", func(t *testing.T) {
		input := `user:
    name: John`

		expected := `user:
    name: Ivan
`

		doc, err := yamlpath.Unmarshal([]byte(input))
		require.NoError(t, err)

		err = doc.Update(yamlpath.NewPointer("user.name"), "Ivan")
		require.NoError(t, err)

		res, err := doc.Marshal()
		require.NoError(t, err)

		assert.Equal(t, expected, string(res))
	})

	t.Run("update map in slice", func(t *testing.T) {
		input := `users:
    - name: John
`

		expected := `users:
    - name: Ivan
`

		doc, err := yamlpath.Unmarshal([]byte(input))
		require.NoError(t, err)

		err = doc.Update(yamlpath.NewPointer("users.0.name"), "Ivan")
		require.NoError(t, err)

		res, err := doc.Marshal()
		require.NoError(t, err)

		assert.Equal(t, expected, string(res))
	})
}
