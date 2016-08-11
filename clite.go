package clite

import "fmt"

type Program struct {
	Name       string
	Commands   map[string]*Command
	longestCmd int
}

type Command struct {
	Program      string
	Name         string
	Summary      string
	Handler      CommandHandler
	Arguments    map[string]*Arg
	longestArg   int
	requiredArgs int
}

type Arg struct {
	Name     string
	Summary  string
	required bool
}

type CommandHandler func([]string)

const DELIM = "    "

func (p *Program) Run(args []string) {
	cmd := args[0]
	cmdObj, ok := p.Commands[cmd]
	if ok {
		if cmdObj.validate(args[1:]) {
			cmdObj.Handler(args[1:])
		} else {
			fmt.Println(cmdObj.String())
		}
	} else {
		fmt.Println(p.String())
	}
}

func NewProgram(name string) *Program {
	return &Program{name, make(map[string]*Command), 0}
}

func (p *Program) HasCommand(name, summary string, handler CommandHandler) *Program {
	cmdLen := len(name)
	cmd := &Command{p.Name, name, summary, handler, make(map[string]*Arg), 0, 0}
	p.Commands[name] = cmd
	if cmdLen > p.longestCmd {
		p.longestCmd = cmdLen
	}
	return p
}

func (p *Program) Command(name string) *Command {
	return p.Commands[name]
}

func (c *Command) HasRequiredArg(name, summary string) *Command {
	return c.addArg(name, summary, true)
}

func (c *Command) HasOptionalArg(name, summary string) *Command {
	return c.addArg(name, summary, false)
}

func (c *Command) addArg(name, summary string, required bool) *Command {
	argLen := len(name)
	arg := &Arg{name, summary, required}
	c.Arguments[name] = arg
	if argLen > c.longestArg {
		c.longestArg = argLen
	}
	if required {
		c.requiredArgs += 1
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

func (c *Command) validate(args []string) bool {

	return len(args) >= c.requiredArgs && len(args) <= len(c.Arguments)

}

func (p *Program) String() string {

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

func (c *Command) String() string {

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
