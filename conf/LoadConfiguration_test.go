package conf

import "testing"

func TestGetConfiguration(t *testing.T) {
	got := GetConfiguration()
	serverOut = Serverout{
		name: "G",
	}
	want := &Config{
		Logsfolderpath: "/lol",
		serverOut,
	}
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
