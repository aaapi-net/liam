**liam** is a *mail* library which provides a set of extensions on go's standard `net/smtp` library,

so all interfaces are just a syntax sugar on the standard ones.

Major additional concepts are:

- *SOE*: Simple, Original, Extended

- `liam.Send()` to quick result

## Usage

### Simple

```go
    package main
    
    import (
        "github.com/aaapi-net/liam"
    )

    err := liam.Send("mail.aaapi.net", 587, "from@aaapi.net", "p@ssw0rd", "to@aaapi.net", "Hello", "World!")
```

### Extended



```go
    package main

    import (
        "github.com/aaapi-net/liam"
    )

	err := liam.
		Smtp("mail.aaapi.net", 587).
		Auth("from@aaapi.net", "p@ssw0rd").
		AddTo("to@aaapi.net").
		Subject("Hello").
                Struct(&Message{Head: "Hello", Body: "World", Footer:"!"}).AsTemplate("{{ .Head }}\r\n{{ .Body }}\r\n{{ .Footer }}")
		Send()
```

```go
package main

import (
	"fmt"
	"github.com/aaapi-net/liam"
)

var mail = liam.
		Smtp("mail.aaapi.net", 587).
		Auth("from@aaapi.net", "p@ssw0rd").
		AddTo("to@aaapi.net")


func main() {
	err := mail.
		Subject("Hello").
		Template(
			map[string]interface{}{
				"Username": "Dave", 
				"Type": "a new client",
			}, 
			`Hello {{ .Username }}, your are {{ .Type }}.`).
		Send()

	fmt.Println(err)
}

```

**Alternatives:**

https://github.com/xhit/go-simple-mail