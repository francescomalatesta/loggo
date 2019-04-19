package commands

import "testing"

func TestItShouldHandleLogCommand(t *testing.T) {
	lc := LogCommand{}

	if !lc.Handles("LOG ERROR OtherArgument") {
		t.Errorf("Error: LogCommand should handle LOG as string starter")
	}
}
func TestItShouldNotHandleOtherCommand(t *testing.T) {
	lc := LogCommand{}

	if lc.Handles("WATLOG ERROR OtherArgument") {
		t.Errorf("Error: LogCommand should only handle LOG as string starter")
	}
}

func TestExecuteShouldReturnOK(t *testing.T) {
	lc := LogCommand{}

	if lc.Execute("LOG ERROR My Error Message") != "OK" {
		t.Errorf("Error: LogCommand should return OK")
	}
}
