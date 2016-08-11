# Usage display example

1. Program usage text

```
> program
usage: program <command>

Available commands are:
    cmd        some random command
    cmd2       another random command
    fooCmd     another one, oh it's another one!

```

2. Usage text format per command

```
> program cmd
usage: program cmd <requiredArg> [<optionalArg>] [<optionalArg2>]

Arguments are:
    requiredArg      some argument that is required to run the command
    optionalArg      some argument that is optional
    optionalArg2     some argument that is also optional

```

The format is heavily influenced by [Hashicorp's cli tool](https://github.com/mitchellh/cli)

# Code example

```go

package main

import (
    "fmt"
    "github.com/shinmyung0/goformat"
)

func main() {
    prog := goformat.Program("program")
    prog.HasCommand("cmd", "some random command")
    prog.HasCommand("cmd2", "another random command")
    prog.HasCommand("foocmd", "another one, it's another one!")

    cmd := prog.Command("cmd")
    cmd.HasArg("required", "some argument that is required to run the command", true)
    cmd.HasArg("optionalArg", "some argument that is optional", false)
    cmd.HasArg("optionalArg2", "some argument that is also optional", false)

    // Will print main usage text
    fmt.Print(format)
    // Will print usage text for cmd
    fmt.Print(format.ForCommand("cmd"))

}

```
