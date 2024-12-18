# yamlpath

yamlpath - go library for working with yaml as path

## Usage

### Get value

```go
package main

import (
	"fmt"
	"github.com/artarts36/yamlpath"
)

const yaml = `
user:
  name: John`

func main() {
	fmt.Print(yamlpath.Get([]byte(yaml), "user.name")) // John
}
```

### Update value

```go
package main

import (
	"fmt"
	"github.com/artarts36/yamlpath"
)

const yaml = `
user:
  name: John`

func main() {
	yamlpath.Update([]byte(yaml), "user.name", "Ivan")
}
```

### Append value to slice

```go
package main

import (
	"fmt"
	"github.com/artarts36/yamlpath"
)

const yaml = `
user:
  phones: [1, 2]`

func main() {
	yamlpath.Append([]byte(yaml), "user.phones", "3")
}
```
