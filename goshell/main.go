package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	args := os.Args
	if len(args) > 1 {
		shell, err := os.Open(args[1])
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}
		scanner := bufio.NewScanner(shell)
		_, _ = fmt.Fprintln(os.Stdout, "welcome to use goshell!")
		for scanner.Scan() { // internally, it advances token based on sperator
			line := scanner.Text()
			if strings.HasPrefix(line, "#") {
				continue
			}
			err = execInput(line)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
			}
		}
		return
	}
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	_ = os.Chdir(usr.HomeDir)
	for {
		showLeader()
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		} // Handle the execution of the input.
		err = execInput(input)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
} // ErrNoPath is returned when 'cd' was called without a second argument.var ErrNoPath = errors.New("path required")

func showLeader() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Print("> ")
		return
	}
	fmt.Print(fmt.Sprintf("%s >", wd))
}

func execInput(input string) error { // Remove the newline character.
	input = strings.TrimSuffix(input, "\n") // Split the input separate the command and the arguments.
	args := strings.Split(input, " ")       // Check for built-in commands.
	switch args[0] {
	case "cd": // 'cd' to home with empty path not yet supported.
		if len(args) < 2 {
			return fmt.Errorf("error path")
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		} // Stop further processing.
		return nil
	case "exit":
		os.Exit(0)
	case "dayin":
		_, _ = fmt.Fprintln(os.Stdout, strings.Join(args[1:], ""))
		return nil
	}
	cmd := exec.Command(args[0], args[1:]...) // Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and save it's output.
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
