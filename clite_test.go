package clite

import "testing"

func handler(args []string) {
	lastCalled = args
}

var prog *Program

var lastCalled []string

func init() {
	prog = NewProgram("program")
	prog.HasCommand("cmd", "some random command", handler)
	prog.HasCommand("cmd2", "another random command", handler)
	prog.HasCommand("fooCmd", "another one, oh it's another one!", handler)

	cmd := prog.Command("cmd")
	cmd.HasRequiredArg("requiredArg", "some argument that is required to run the command")
	cmd.HasOptionalArg("optionalArg", "some argument that is optional")
	cmd.HasOptionalArg("optionalArg2", "some argument that is also optional")
}

func TestProgramFormat(t *testing.T) {

	// Exact length the string should be
	exactLen := 169
	if len(prog.String()) != exactLen {
		t.Errorf("The formatted string length is incorrect. Should be %d\n", exactLen)
	}
}

func TestCommandFormat(t *testing.T) {

	cmd := prog.Command("cmd")
	exactLen := 259
	if len(cmd.String()) != exactLen {
		t.Errorf("The formatted string length is incorrect. Should be %d\n", exactLen)
	}

}

func TestRunCommand(t *testing.T) {
	prog.Run([]string{"cmd", "testing"})
	if len(lastCalled) != 1 {
		t.Error("Wrong number of arguments detected")
	}
}

func TestRunCommandFailed(t *testing.T) {
	prog.Run([]string{"noncmd", "testing", "2ndarg"})
	if len(lastCalled) == 2 {
		t.Error("Handler called for non existant command")
	}
}
