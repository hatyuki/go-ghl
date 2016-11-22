package ghl

import (
	"net/http"
	"regexp"
	"runtime"
	"strings"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// Version of this program
const Version = "0.0.1"

type GHL struct {
	Client *github.Client
}

func NewClient(token string) (g *GHL) {
	var tc *http.Client

	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc = oauth2.NewClient(oauth2.NoContext, ts)
	}

	g = &GHL{Client: github.NewClient(tc)}

	return
}

func (g *GHL) GetDownloadURL(target, os, arch string) (location string, err error) {
	if target == "" {
		err = errors.New("target is required")
		return
	}

	if os == "" {
		os = runtime.GOOS
	}

	if arch == "" {
		arch = runtime.GOARCH
	}

	owner, repo, tag, err := getOwnerRepoAndTag(target)
	if err != nil {
		err = errors.Wrap(err, "failed to resolve target")
		return
	}

	var release *github.RepositoryRelease

	if tag == "" {
		release, _, err = g.Client.Repositories.GetLatestRelease(owner, repo)
	} else {
		release, _, err = g.Client.Repositories.GetReleaseByTag(owner, repo, tag)
	}
	if err != nil {
		err = errors.Wrap(err, "failed to fetch a release")
		return
	}

	for _, asset := range release.Assets {
		if strings.Contains(*asset.Name, os) && strings.Contains(*asset.Name, arch) {
			location = *asset.BrowserDownloadURL
			break
		}
	}
	if location == "" {
		errors.New("no release assets available")
		return
	}

	return
}

var targetReg = regexp.MustCompile(`^([^/]+)/([^@]+)(?:@(.+))?$`)

func getOwnerRepoAndTag(target string) (owner, repo, tag string, err error) {
	matches := targetReg.FindStringSubmatch(target)

	if len(matches) != 4 {
		err = errors.New("failed to get owner, repo and tag")
		return
	}

	owner = matches[1]
	repo = matches[2]
	tag = matches[3]

	return
}
