package conf

import "testing"

func TestGetConfiguration(t *testing.T) {
	got, _ := GetConfiguration("../config.json")
	serverOut := ServerOut{
		Name: "serverOut_",
	}
	want := &Config{
		Logsfolderpath: "./LOGS/",
		ServerOut:      serverOut,
	}
	if *got != *want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
