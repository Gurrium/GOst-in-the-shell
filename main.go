package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := os.Stdout
	stderr := os.Stderr
	gsh_loop(stdin, stdout, stderr)
}

func gsh_loop(stdin *bufio.Reader, stdout io.Writer, stderr io.Writer) {
	for {
		fmt.Fprint(stdout, "> ")
		line, err := gsh_read_line(stdin)
		if err != nil {
			fmt.Fprint(stderr, "Couldn't read your input")
		}

		name, args := gsh_split_line(line)
		gsh_execute(stdout, stderr, name, args)
	}
}

func gsh_read_line(stdin *bufio.Reader) (line string, err error) {
	return stdin.ReadString('\n')
}

func gsh_split_line(line string) (name string, args []string) {
	trimmedLine := bytes.TrimSpace([]byte(line))
	result := strings.Split(string(trimmedLine), " ")
	if len(result) > 1 {
		return result[0], result[1:len(result)]
	} else {
		return result[0], nil
	}
}

func gsh_execute(stdout io.Writer, stderr io.Writer, name string, args []string) {
	cmd := exec.Command(name, args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(stderr, err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(stderr, err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprintln(stderr, err)
		return
	}

	out, _ := ioutil.ReadAll(stdoutPipe)
	fmt.Fprint(stdout, string(out))
	slurp, _ := ioutil.ReadAll(stderrPipe)
	if len(slurp) > 0 {
		fmt.Fprintln(stderr, string(slurp))
	}

	if err := cmd.Wait(); err != nil {
		fmt.Fprintln(stderr, err)
	}
}
