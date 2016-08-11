package goformat

import "fmt"

type ProgramFormat struct {
	Name       string
	Commands   map[string]*CommandFormat
	longestCmd int
}

type CommandFormat struct {
	Program    string
	Name       string
	Summary    string
	Arguments  map[string]*Arg
	longestArg int
}

type Arg struct {
	Name     string
	Summary  string
	required bool
}

const DELIM = "    "

func Program(name string) *ProgramFormat {
	return &ProgramFormat{name, make(map[string]*CommandFormat), 0}
}

func (p *ProgramFormat) Command(name string) *CommandFormat {
	return p.Commands[name]
}

func (p *ProgramFormat) HasCommand(name, summary string) *ProgramFormat {
	cmdLen := len(name)
	cmd := &CommandFormat{p.Name, name, summary, make(map[string]*Arg), 0}
	p.Commands[name] = cmd
	if cmdLen > p.longestCmd {
		p.longestCmd = cmdLen
	}
	return p
}

func (c *CommandFormat) HasArg(name, summary string, required bool) *CommandFormat {
	argLen := len(name)
	arg := &Arg{name, summary, required}
	c.Arguments[name] = arg
	if argLen > c.longestArg {
		c.longestArg = argLen
	}
	return c
}

func genPadding(num int) string {
	result := ""
	for ; num > 0; num-- {
		result += " "
	}
	return result
}

func (p *ProgramFormat) String() string {

	result := fmt.Sprintf("usage: %s <command>\n\n", p.Name)
	result += "Available commands are:\n"
	for cmd, frmt := range p.Commands {
		result += DELIM
		result += cmd
		result += genPadding(p.longestCmd - len(cmd))
		result += DELIM
		result += frmt.Summary
		result += "\n"
	}
	return result
}

func (c *CommandFormat) String() string {

	result := fmt.Sprintf("usage: %s %s", c.Program, c.Name)
	for arg, frmt := range c.Arguments {
		result += " "
		if !frmt.required {
			result += "[<" + arg + ">]"
		} else {
			result += "<" + arg + ">"
		}
	}
	if len(c.Arguments) > 0 {
		result += "\n\n"
		result += "Arguments are:\n"
		for arg, frmt := range c.Arguments {
			result += DELIM
			result += arg
			result += genPadding(c.longestArg - len(arg))
			result += DELIM
			result += frmt.Summary
			result += "\n"
		}
	}

	return result
}
