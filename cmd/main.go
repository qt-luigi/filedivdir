package main

import (
	"fmt"
	"os"

	"github.com/qt-luigi/filedivdir"
)

const (
	usage = `filedivdir divides the files in the base directory by creating a directory for each file timestamp.

Usage:

	filedivdir [argument]

The argument is:

	basedir    the base directory that contains the files to be divide.
`
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	if err := filedivdir.Divide(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
