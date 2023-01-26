package main

import (
	"fmt"
	"strings"
)

type cmd1 struct {
	Path string
	Args []string
}

type runner interface {
	Run() error
	Help() string
}

func (c *cmd1) Run() error {
	fmt.Printf("Running %s with Args %s", c.Path, strings.Join(c.Args, " "))
	return nil
}

func (c *cmd1) Help() string {
	return fmt.Sprintf("Usage: %s ", c.Path)
}

func runCmd(r runner) error {
	return r.Run()
}

func main() {

	runCmd(&cmd1{
		Path: "/bin/ls",
		Args: []string{"/tmp", "/var/log"}})

}
