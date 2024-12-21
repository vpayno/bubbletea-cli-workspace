package clis

import "testing"

func TestRoot(_ *testing.T) {
	/*
	  This is an empty test in the root package.
	  Without it `go test -v .` prints "testing: warning: no tests to run".
	  Or you get `?		  github.com/org_name/repo_name  [no test files]` when with `./...`.
	*/
}
