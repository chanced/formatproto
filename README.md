# formatproto

formatproto is a post-processor for [protoc-gen-star](https://github.com/lyft/protoc-gen-star) that formats proto files with [clang-format](https://clang.llvm.org/docs/ClangFormatStyleOptions.html)

## Usage:

```go
package main

import (
    "github.com/chanced/formatproto"
    pgs "github.com/lyft/protoc-gen-star"
)

func main() {
    pgs.Init(
        pgs.DebugEnv("DEBUG"),
    ).RegisterModule(
        // your modules here
    ).RegisterPostProcessor(
        formatproto.PostProcessor(),
    ).Render()
}
```

You will also need a `.clang-format` file at the root of your directory. An example file:

```yaml
Language: Proto
BasedOnStyle: google
IndentWidth: 2
```

## Prereqs:

You must have clang-format installed.

### MacOS

`brew install clang-format`

### Linux (with apt)

`sudo apt install clang-format`

### Windows

Download the binary: [https://llvm.org/builds/](https://llvm.org/builds/)

## License

MIT
