package goformat

import (
	"fmt"
	"testing"
)

func TestProgramFormat(t *testing.T) {

	prog := Program("program")
	prog.HasCommand("cmd", "some random command")
	prog.HasCommand("cmd2", "another random command")
	prog.HasCommand("fooCmd", "another one, oh it's another one!")

	// Exact length the string should be
	exactLen := 169
	if len(prog.String()) != exactLen {
		t.Errorf("The formatted string length is incorrect. Should be %d\n", exactLen)
	} else {
		fmt.Println(prog)

	}

}

func TestCommandFormat(t *testing.T) {

	prog := Program("program")
	prog.HasCommand("cmd", "some random command")
	prog.HasCommand("cmd2", "another random command")
	prog.HasCommand("foocmd", "another one, oh it's another one!")
	cmd := prog.Command("cmd")
	cmd.HasArg("requiredArg", "some argument that is required to run the command", true)
	cmd.HasArg("optionalArg", "some argument that is optional", false)
	cmd.HasArg("optionalArg2", "some argument that is also optional", false)

	exactLen := 259
	if len(cmd.String()) != exactLen {
		t.Errorf("The formatted string length is incorrect. Should be %d\n", exactLen)
	} else {
		fmt.Println(cmd.String())

	}

}
