# clite

Lightweight cli library for golang. You just declare what commands and arguments you have in a simple syntax.
It doesn't handle flags (yet!), but personally I don't like them anyways because often times
flags just aren't very descriptive.

Mainly used for setting up simple clis with formatted usage displays, and hooking up simple handlers.
More complex flag checking / argument validation might be part of it, but I feel like that should be
part of the handler function's job anyways.

# Code example

```go

package main

import (
    "fmt"
    "os"
    "github.com/shinmyung0/clite"
)

func main() {
    prog := clite.NewProgram("program")
    prog.HasCommand("cmd", "some random command", HandlerFn)
    prog.HasCommand("cmd2", "another random command", HandlerFm)
    prog.HasCommand("fooCmd", "another one, it's another one!", HandlerFn)

    cmd := prog.Command("cmd")
    // This is mainly for setting up the command specific usage texts
    // As well as simple validation that the user passes in numRequired <= len(args) <= numTotal arguments
    cmd.HasRequiredArg("required", "some argument that is required to run the command")
    cmd.HasOptionalArg("optionalArg", "some argument that is optional")
    cmd.HasOptionalArg("optionalArg2", "some argument that is also optional")

    // Will print main usage text
    fmt.Println(format)
    // Will print usage text for cmd
    fmt.Println(format.ForCommand("cmd"))

    args := os.Args[1:]
    // Assumes that the slice doesn't include the execution path
    // also returns an exit code as an int, which should be returned by the handlers
    os.Exit(prog.Run(args))

}

func HandlerFn(args []string) int {
    return 0
}

```

# Usage display example

1. Main program usage text

```
> program
usage: program <command>

Available commands are:
    cmd        some random command
    cmd2       another random command
    fooCmd     another one, oh it's another one!

```

2. Command specific usage text

```
> program cmd
usage: program cmd <requiredArg> [<optionalArg>] [<optionalArg2>]

Arguments are:
    requiredArg      some argument that is required to run the command
    optionalArg      some argument that is optional
    optionalArg2     some argument that is also optional

```

The format is heavily influenced by [Hashicorp's cli tool](https://github.com/mitchellh/cli)



# Possible Todo
- Add flags (maybe!)
- Add dependent arguments (ie. an argument that is optional but only to be used in conjunction with another argument)
