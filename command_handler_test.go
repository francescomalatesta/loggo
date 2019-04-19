package main

import (
	"testing"

	"github.com/francescomalatesta/loggo/commands"
)

func TestItShouldHandleRecognizedCommand(t *testing.T) {
	cm := map[string]commands.Command{
		"hello": commands.HelloCommand{},
	}

	result, err := HandleInputCommandAndReturnResult("HELLO", cm)

	if result != "Hello to you." && err != nil {
		t.Errorf("Command handler is not recognizing commands right")
	}
}

func TestItShouldHandleUnrecognizedCommand(t *testing.T) {
	cm := map[string]commands.Command{
		"hello": commands.HelloCommand{},
	}

	_, err := HandleInputCommandAndReturnResult("OTHER", cm)

	if err == nil {
		t.Errorf("Command handler function should throw an error")
	}
}
