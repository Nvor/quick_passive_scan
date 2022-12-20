package test

import (
	"testing"
	"quick_passive_recon/pkg/args_proc"
)

func TestProcessArgsHasArgs(t *testing.T) {
	want := "example.com"
	cmd := args_proc.ProcessArgs([]string {"example.com", "fast"})

	if cmd.Url == "" {
		t.Errorf("got %q, wanted %q", cmd.Url, want)
	}
	if cmd.Mode == "" {
		t.Errorf("got %q, wanted %q", cmd.Mode, want)
	}
}

func TestProcessUrl(t *testing.T) {
	cmd := args_proc.Command{}
	args_proc.ProcessUrl(&cmd, "example.com")
	want := "example.com"

	if cmd.Url != want {
		t.Errorf("got %q, wanted %q", cmd.Url, want)
	}
}

func TestProcessMode(t *testing.T) {
	cmd := args_proc.Command{}
	args_proc.ProcessMode(&cmd, "fast")
	want := "fast"

	if cmd.Mode != want {
		t.Errorf("got %q, wanted %q", cmd.Mode, want)
	}
}