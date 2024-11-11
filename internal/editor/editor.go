package editor

import (
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/term"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	defaultEditor = "vi"
	defaultShell  = "/bin/bash"
	windowsEditor = "notepad"
	windowsShell  = "cmd"
)

type Editor struct {
	Args  []string
	Shell bool
}

func (e Editor) LaunchTempFile(prefix, suffix string, r io.Reader) ([]byte, string, error) {
	file, err := os.CreateTemp("", prefix+suffix)
	if err != nil {
		return nil, "", err
	}
	path := file.Name()
	internal.VerboseLog("temp file created: %s", path)
	if _, err := io.Copy(file, r); err != nil {
		_ = os.Remove(path)
		return nil, "", err
	}
	_ = file.Close()
	if err := e.Launch(path); err != nil {
		return nil, path, err
	}
	bytes, err := os.ReadFile(path)
	return bytes, path, err
}

func (e Editor) Launch(path string) error {
	if len(e.Args) == 0 {
		return fmt.Errorf("no editor defined, can't open %s", path)
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	args := e.args(abs)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	fmt.Printf("Open file with editor: %v\n", args)
	if err := (term.TTY{In: os.Stdin, TryDev: true}).Safe(cmd.Run); err != nil {
		if err, ok := err.(*exec.Error); ok {
			if err.Err == exec.ErrNotFound {
				return fmt.Errorf("unable to launch the editor %q", strings.Join(e.Args, " "))
			}
		}
		return fmt.Errorf("there was a problem with the editor %q", strings.Join(e.Args, " "))
	}
	return nil
}

func (e Editor) args(path string) []string {
	args := make([]string, len(e.Args))
	copy(args, e.Args)
	if e.Shell {
		last := args[len(args)-1]
		args[len(args)-1] = fmt.Sprintf("%s %q", last, path)
	} else {
		args = append(args, path)
	}
	return args
}

func NewDefaultEditor(envs []string) Editor {
	args, shell := defaultEnvEditor(envs)
	return Editor{
		Args:  args,
		Shell: shell,
	}
}

func defaultEnvShell() []string {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = platform(defaultShell, windowsShell)
	}
	flag := "-c"
	if shell == windowsShell {
		flag = "/C"
	}
	return []string{
		shell,
		flag,
	}
}

func defaultEnvEditor(envs []string) ([]string, bool) {
	var editor string
	for _, env := range envs {
		if len(env) > 0 {
			editor = os.Getenv(env)
		}
		if len(editor) > 0 {
			break
		}
	}
	if len(editor) == 0 {
		editor = platform(defaultEditor, windowsEditor)
	}
	if !strings.Contains(editor, " ") {
		return []string{editor}, false
	}
	if !strings.ContainsAny(editor, "\"'\\") {
		return strings.Split(editor, " "), false
	}
	shell := defaultEnvShell()
	return append(shell, editor), true
}

func platform(linux, windows string) string {
	if runtime.GOOS == "windows" {
		return windows
	}
	return linux
}
