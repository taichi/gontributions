package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var name, email, repo, msg string
var commits int
var None struct{}

func init() {
	flag.StringVar(&name, "name", "", "author and commiter name")
	flag.StringVar(&email, "email", "", "author and commiter email address")
	flag.StringVar(&repo, "repo", "./repo", "new repository path")
	flag.IntVar(&commits, "base", 10, "number of commits for graph")
	flag.StringVar(&msg, "msg", "happy hacking !!", "commit message")
}

func main() {
	flag.Parse()
	source := parse()

	willBeExit(None, os.MkdirAll(repo, os.ModeDir))
	run("git", []string{"init"})

	oneday := 24 * time.Hour
	date := time.Now().Add(-1 * oneday * 366)

	opt := []string{"commit", "--allow-empty", "-m", msg}
	env := makeEnv()

	for _, value := range source {
		if 0 < value {
			d := date.Format(time.RFC1123Z)
			env["GIT_AUTHOR_DATE"] = d
			env["GIT_COMMITTER_DATE"] = d
			times := value * commits
			fmt.Printf("%s [%3d]\n", d, times)
			for i := 0; i < times; i++ {
				run("git", opt, env)
			}
		}
		date = date.Add(oneday)
	}
}

func parse() []int {
	args := flag.Args()
	path := "source.txt"
	if 0 < len(args) {
		path = args[0]
	}

	file, ferr := os.Open(path)
	willBeExit(None, ferr)
	defer file.Close()

	data := make([]int, 350)
	for s, ln := bufio.NewScanner(file), 0; s.Scan() && ln < 7; ln++ {
		line := s.Text()
		if 50 < len(line) {
			line = line[0:50]
		} else {
			line = fmt.Sprintf("%-50s", line)
		}
		for i, r := range line {
			pos := i*7 + ln
			data[pos] = multiplyer(r)
		}
	}
	return append([]int{1, 1, 1, 1, 1, 1}, data...)
}

func multiplyer(r rune) int {
	switch r {
	case '1':
		fallthrough
	case '-':
		return 1
	case '2':
		fallthrough
	case '*':
		return 2
	case '3':
		fallthrough
	case '@':
		return 3
	case '4':
		fallthrough
	case '#':
		return 4
	}
	return 0
}

func run(bin string, options []string, env ...map[string]string) {
	willBeExit(exec.LookPath(bin))

	cmd := exec.Command(bin, options...)
	cmd.Dir = repo
	cmd.Env = mergeEnv(env...)

	willBeExit(None, cmd.Run())
}

func makeEnv() map[string]string {
	env := map[string]string{}
	if 0 < len(name) {
		env["GIT_AUTHOR_NAME"] = name
		env["GIT_COMMITTER_NAME"] = name
	}
	if 0 < len(email) {
		env["GIT_AUTHOR_EMAIL"] = email
		env["GIT_COMMITTER_EMAIL"] = email
	}
	return env
}

func mergeEnv(newmaps ...map[string]string) []string {
	in := []string{}
	for _, m := range newmaps {
		for k, v := range m {
			in = append(in, fmt.Sprintf("%s=%s", k, v))
		}
	}

	out := os.Environ()
NextVar:
	for _, inkv := range in {
		k := strings.SplitAfterN(inkv, "=", 2)[0]
		for i, outkv := range out {
			if strings.HasPrefix(outkv, k) {
				out[i] = inkv
				continue NextVar
			}
		}
		out = append(out, inkv)
	}
	return out
}

func willBeExit(any interface{}, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
