package ghl

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

const (
	EnvGithubToken = "GITHUB_TOKEN"
)

func TestGHL_GetDownloadURL(t *testing.T) {
	token := os.Getenv(EnvGithubToken)
	if token == "" {
		t.Skipf("environment variable %s not set", EnvGithubToken)
	}

	g := NewClient(token)

	if got, err := g.GetDownloadURL("", "", ""); err == nil {
		t.Error("should return an error")
	} else if got != "" {
		t.Errorf("returns wrong URL: got %v want %v", got, "")
	}

	if got, err := g.GetDownloadURL("a/b/c", "", ""); err == nil {
		t.Error("should return an error")
	} else if got != "" {
		t.Errorf("returns wrong URL: got %v want %v", got, "")
	}

	if got, err := g.GetDownloadURL("hatyuki/ghl@foo", "", ""); err == nil {
		t.Error("should return an error")
	} else if got != "" {
		t.Errorf("returns wrong URL: got %v want %v", got, "")
	}

	if got, err := g.GetDownloadURL("hatyuki/ghl", "", ""); err != nil {
		t.Errorf("returns an error: got %v", err)
	} else if got == "" {
		t.Error("should return URL")
	} else if strings.Contains(got, runtime.GOOS) && strings.Contains(got, runtime.GOARCH) {
		t.Errorf("returns wrong URL: got %v", got)
	}

	os := "linux"
	arch := "386"
	if got, err := g.GetDownloadURL("hatyuki/ghl", os, arch); err != nil {
		t.Errorf("returns an error: got %v", err)
	} else if got == "" {
		t.Error("should return URL")
	} else if strings.Contains(got, os) && strings.Contains(got, arch) {
		t.Errorf("returns wrong: got %v", got)
	}

	if got, err := g.GetDownloadURL("hatyuki/ghl@0.0.1", "", ""); err != nil {
		t.Errorf("returns an error: got %v", err)
	} else if got == "" {
		t.Error("should return URL")
	} else if tag := "0.0.1"; strings.Contains(got, tag) {
		t.Errorf("returns wrong URL: got %v", got)
	}
}
