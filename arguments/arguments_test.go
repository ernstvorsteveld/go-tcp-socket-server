package arguments

import "testing"

func Test_handle_valid_arguments(t *testing.T) {
	expectedServerName := "test server"
	expectedPortNumer := "8080"
	args := []string{"program", expectedServerName, expectedPortNumer}
	actual := HandleArguments(args)
	if actual.Server != expectedServerName {
		t.Errorf("Incorrect server name found: %s, expected %s.", actual.Server, expectedServerName)
	}
	if actual.Port != expectedPortNumer {
		t.Errorf("Incorrect port number found: %s, expected %s.", actual.Port, expectedPortNumer)
	}
}
