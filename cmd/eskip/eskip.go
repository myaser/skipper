package main

import (
	"errors"
	"fmt"
	"os"
)

type (
	command     string
	commandFunc func(in, out *medium) error
)

const (
	check  command = "check"
	print  command = "print"
	upsert command = "upsert"
	reset  command = "reset"
	delete command = "delete"
	help   command = "help"
)

var commands = map[command]commandFunc{
	check:  checkCmd,
	print:  printCmd,
	upsert: upsertCmd,
	reset:  resetCmd,
	delete: deleteCmd,
	help:   helpCmd}

var (
	missingCommand = errors.New("missing command")
	invalidCommand = errors.New("invalid command")
)

func printStderr(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

func exit(err error) {
	if err == nil {
		os.Exit(0)
	}

	printStderr(err)
	printStderr()
	hint()
	os.Exit(-1)
}

func getCommand() (command, error) {
	if len(os.Args) < 2 {
		return "", missingCommand
	}

	cmd := command(os.Args[1])
	if _, ok := commands[cmd]; ok {
		return cmd, nil
	} else {
		return "", invalidCommand
	}
}

func upsertCmd(in, out *medium) error {
	return nil
}

func resetCmd(in, out *medium) error {
	return nil
}

func deleteCmd(in, out *medium) error {
	return nil
}

func main() {
	cmd, err := getCommand()
	if err != nil {
		exit(err)
	}

	if cmd == help {
		exit(helpCmd(nil, nil))
	}

	media, err := processArgs()
	if err != nil {
		exit(err)
	}

	in, out, err := validateSelectMedia(cmd, media)
	if err != nil {
		exit(err)
	}

	exit(commands[cmd](in, out))
}
