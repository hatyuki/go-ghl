package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/hatyuki/go-ghl"
	"github.com/jessevdk/go-flags"
	"github.com/tcnksm/go-gitconfig"
)

const (
	notExit = iota - 1
	exitOK
	exitErr
	exitParseArgsErr

	// EnvGitHubToken is environmental var to set GitHub API token
	EnvGithubToken = "GITHUB_TOKEN"
)

var build = "devel"

type options struct {
	OS    string `long:"os" description:"specifies the name of the target OS like 'windows', 'darwin' or 'linux'."`
	Arch  string `long:"arch" description:"specifies the name of the target architecture like '386' or 'amd64'."`
	Token string `long:"token" description:"GitHub API Token."`
}

func main() {
	os.Exit(Run(os.Args[1:]))
}

// Run invokes the CLI with the given arguments.
func Run(args []string) int {
	target, opts, status := parseOptions(args)
	if status != notExit {
		return status
	}

	if opts.Token == "" {
		opts.Token = getGithubToken()
	}

	client := ghl.NewClient(opts.Token)
	location, err := client.GetDownloadURL(target, opts.OS, opts.Arch)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	fmt.Println(location)

	return exitOK
}

func parseOptions(args []string) (target string, opts *options, status int) {
	opts = &options{OS: runtime.GOOS, Arch: runtime.GOARCH}
	parser := flags.NewParser(opts, flags.PrintErrors|flags.PassDoubleDash)
	parser.Usage = fmt.Sprintf(`[OPTIONS] <owner>/<repo>[@<tag>]

Version:
  %s (build: %s)`, ghl.Version, build)

	remains, err := parser.ParseArgs(args)
	if err != nil {
		parser.WriteHelp(os.Stderr)
		status = exitParseArgsErr
		return
	}

	if len(remains) == 0 {
		parser.WriteHelp(os.Stderr)
		status = exitOK
		return
	}

	target = remains[0]
	status = notExit

	return
}

func getGithubToken() (token string) {
	token = os.Getenv(EnvGithubToken)
	if token != "" {
		return
	}

	token, _ = gitconfig.GithubToken()
	return
}
