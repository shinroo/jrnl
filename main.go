package main

import (
	"fmt"
	"os"

	"github.com/andrewpillar/cli"

	"github.com/andrewpillar/jrnl/command"
)

func main() {
	c := cli.New()

	c.Main(command.Main)

	c.AddFlag(&cli.Flag{
		Name: "help",
		Long: "--help",
	})

	c.Command("init", command.Initialize)
	c.Command("tmpl", command.Template)

	postCmd := c.Command("post", command.Post)

	postCmd.AddFlag(&cli.Flag{
		Name:     "category",
		Short:    "-c",
		Long:     "--category",
		Argument: true,
		Default:  "",
	})

	c.Command("edit", command.ChangePost)
	c.Command("rm", command.ChangePost)

	listCmd := c.Command("ls", command.List)

	listCmd.AddFlag(&cli.Flag{
		Name:     "category",
		Short:    "-c",
		Long:     "--category",
		Argument: true,
		Default:  "",
	})

	remoteCmd := c.Command("remote", command.Remote)

	remoteCmd.Command("ls", command.RemoteList)

	remoteSetCmd := remoteCmd.Command("set", command.RemoteSet)

	remoteSetCmd.AddFlag(&cli.Flag{
		Name:  "default",
		Short: "-d",
		Long:  "--default",
	})

	remoteSetCmd.AddFlag(&cli.Flag{
		Name:     "port",
		Short:    "-p",
		Long:     "--port",
		Argument: true,
		Default:  22,
	})

	remoteCmd.Command("rm", command.RemoteRemove)

	publishCmd := c.Command("publish", command.Publish)

	publishCmd.AddFlag(&cli.Flag{
		Name:  "draft",
		Short: "-d",
		Long:  "--draft",
	})

	publishCmd.AddFlag(&cli.Flag{
		Name:  "remote",
		Short: "-r",
		Long:  "--remote",
	})

	if err := c.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
