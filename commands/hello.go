package commands

type HelloCommand struct{}

func (hc HelloCommand) Handles(s string) bool {
	return s == "HELLO"
}

func (hc HelloCommand) GetResponse(s string) string {
	return "Hello to you."
}
