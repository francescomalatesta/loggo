package commands

type Command interface {
	Handles(s string) bool
	GetResponse(s string) string
}
