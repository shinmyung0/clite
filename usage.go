package formatter

type UsageFormatter struct {
	Program  string
	Commands map[string]Options
}

type Options struct {
	Args map[string]Arg
}

type Arg struct {
	Summary  string
	Optional bool
}

const DELIM = "    "

func NewFormatter(cmd string) UsageFormatter {
	return UsageFormatter{cmd, make(map[string]Options)}
}

func (u *UsageFormatter) AddCommand(cmd, summary string, optional bool) {

}

func (c *UsageFormatter) PrintUsage() {

}
