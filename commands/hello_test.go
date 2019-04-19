package commands

import "testing"

func TestItShouldHandleCommandWithRightTextCommand(t *testing.T) {
	c := HelloCommand{}

	if !c.Handles("HELLO") {
		t.Errorf("Error: HelloCommand should handle HELLO")
	}

	if c.Handles("OTHER") {
		t.Errorf("Error: HelloCommand should not handle OTHER")
	}
}

func TestItShouldReturnHelloMessage(t *testing.T) {
	c := HelloCommand{}

	if c.GetResponse("HELLO") != "Hello to you." {
		t.Errorf("Error: HelloCommand.GetResponse not returning the right message")
	}
}
