package commands

// HelloCommand built for learning purposes
type HelloCommand struct{}

// Handles the incoming string to check condition
func (hc HelloCommand) Handles(s string) bool {
	return s == "HELLO"
}

// Execute the command and return a hello message
func (hc HelloCommand) Execute(s string) string {
	return "Hello to you."
}
