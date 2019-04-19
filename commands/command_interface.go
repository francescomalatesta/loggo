package commands

// Command is the standard interface for recognizable commands
type Command interface {
	Handles(s string) bool
	Execute(s string) string
}
