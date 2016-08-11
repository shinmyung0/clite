package goformat

type UsageFormat struct {
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
