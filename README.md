# envconfig

```Go
import "github.com/namsral/envconfig"
```

## Documentation

See [godoc](http://godoc.org/github.com/namsral/envconfig)

## Usage

Set some environment variables:

```Bash
export DEBUG=false
export PORT=8080
export USER=Kelsey
export RATE="0.5"
```

Write some code:

```Go
package main

import (
    "fmt"
    "log"

    "github.com/namsral/envconfig"
)

type Specification struct {
    Debug bool
    Port  int
    User  string
    Rate  float32
}

func main() {
    var s Specification
    err := envconfig.Process(&s)
    if err != nil {
        log.Fatal(err.Error())
    }
    format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\n"
    _, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate)
    if err != nil {
        log.Fatal(err.Error())
    }
}
```

Results:

```Bash
Debug: false
Port: 8080
User: Kelsey
Rate: 0.500000
```

## Struct Tag Support

Envconfig supports the use of struct tags to specify alternate
environment variables.

For example, consider the following struct:

```Go
type Specification struct {
    MultiWordVar string `env:"multi_word_var"`
}
```

Whereas before, the value for `MultiWordVar` would have been populated
with `MULTIWORDVAR`, it will now be populated with
`MULTI_WORD_VAR`.

```Bash
export MULTI_WORD_VAR="this will be the value"

# export MULTIWORDVAR="and this will not"
```

## Mandatory vs Optional Fields

Specification fields are mandatory by default. To make a field optional add the struct tag option `optional`.

For example, the field `Password` in the following struct is optional:

```Go
type Specification struct {
    User string
    Password string `env:",optional"`
}
```