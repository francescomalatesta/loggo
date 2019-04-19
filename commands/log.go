package commands

import (
	"fmt"
	"strings"
	"time"
)

// LogCommand used to log a specific event
type LogCommand struct{}

// Handles the incoming string to check condition
func (lc LogCommand) Handles(s string) bool {
	return strings.HasPrefix(s, "LOG")
}

// Execute the command given a specific string. Returns OK if everything went well
func (lc LogCommand) Execute(s string) string {
	commandAsArray := strings.Split(s, " ")

	severity := commandAsArray[1]
	content := strings.Join(commandAsArray[2:], " ")

	line := time.Now().Format("2006-01-02 15:04:05 MST") + " | " + severity + " | " + content
	fmt.Println(line)

	return "OK"
}
