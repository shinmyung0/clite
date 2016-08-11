# Usage display example

Main program usage text

```bash
> program
usage: program <command>

Available commands are:
    cmd        some random command
    cmd2       another random command
    foocmd     another one, it's another one!

```

Usage text format per command

```bash
> program cmd
usage: program cmd <requiredArg> [<optionalArg>] [<optionalArg2>]

Arguments are:
    requiredArg      some argument that is required to run the command
    optionalArg      some argument that is optional
    optionalArg2     some argument that is also optional

```

# Code example

```go

package main

import (
    "fmt"
    "github.com/shinmyung0/goformat"
)

func main() {
    format := goformat.Program("program")
                      .hasCommand("cmd", "some random command")
                      .hasCommand("cmd2", "another random command")
                      .hasCommand("foocmd", "another one, it's another one!")
    format.Command("cmd")
          .hasArg("required", "some argument that is required to run the command", true)
          .hasArg("optionalArg", "some argument that is optional", false)
          .hasArg("optionalArg2", "some argument that is also optional", false)

    // Will print main usage text
    fmt.Print(format)
    // Will print usage text for cmd
    fmt.Print(format.ForCommand("cmd"))


}


```
