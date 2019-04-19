package main

import (
	"errors"

	"github.com/francescomalatesta/loggo/commands"
)

func HandleInputCommandAndReturnResult(s string, cm map[string]commands.Command) (string, error) {
	for _, command := range cm {
		if command.Handles(s) {
			result := command.GetResponse(s)
			return result, nil
		}
	}

	return "", errors.New("Command not recognized: " + s)
}
