package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Verbose bool

func Log(format string, args ...interface{}) {
	if Verbose {
		_, _ = fmt.Fprintf(os.Stdout, format+"\n", args...)
	}
}

func Error(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func CheckErr(err error) {
	cobra.CheckErr(err)
}
