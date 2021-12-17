package app

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"strings"
)

type Command struct {
	usage    string
	desc     string
	options  CliFlags
	commands []*Command
	runFunc  RunCommandFunc
}

type CommandOption func(*Command)

func WithCommandOptions(opt CliFlags) CommandOption {
	return func(c *Command) {
		c.options = opt
	}
}

type RunCommandFunc func(args []string) error

func WithCommandRunFunc(rcf RunCommandFunc) CommandOption {
	return func(c *Command) {
		c.runFunc = rcf
	}
}

// NewCommand 用于生成 Command
func NewCommand(usage string, desc string, opts ...CommandOption) *Command {
	c := &Command{
		usage: usage,
		desc:  desc,
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// AddCommands AddCommand 用于追加 command
func (c Command) AddCommands(cmd ...*Command) {
	c.commands = append(c.commands, cmd...)
}

func (c *Command) cobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   c.usage,
		Short: c.desc,
	}
	cmd.SetOut(os.Stdout)
	cmd.Flags().SortFlags = false
	if len(c.commands) > 0 {
		for _, command := range c.commands {
			cmd.AddCommand(command.cobraCommand())
		}
	}
	if c.runFunc != nil {
		cmd.Run = c.runCommand
	}
	if c.options != nil {
		for _, f := range c.options.Flags().FlagSets {
			cmd.Flags().AddFlagSet(f)
		}
		// c.flags.AddFlags(cmd.Flags())
	}
	addHelpCommandFlag(c.usage, cmd.Flags())

	return cmd
}

func (c *Command) runCommand(cmd *cobra.Command, args []string) {
	if c.runFunc != nil {
		if err := c.runFunc(args); err != nil {
			fmt.Printf("%v %v\n", color.RedString("Error:"), err)
			os.Exit(1)
		}
	}
}

// FormatBaseName 用于转换应用的文件名称
func FormatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}
