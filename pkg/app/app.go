/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package app

import (
	"fmt"
	"os"

	cliflag "blog-api/pkg/cli/flag"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	progressMessage = color.GreenString("============>")
)

type App struct {
	use   string // 应用名称
	short string
	long  string
	// options
	options   CliOptions
	cmd       *cobra.Command
	commands  []*Command
	args      cobra.PositionalArgs
	runFunc   RunFunc
	silence   bool
	useConfig bool // Use config file
}

type Option func(*App)

type RunFunc func(basename string) error

func WithLong(desc string) Option {
	return func(app *App) {
		app.long = desc
	}
}

func WithOptions(opt CliOptions) Option {
	return func(app *App) {
		app.options = opt
	}
}

func WithRunFunc(runFunc RunFunc) Option {
	return func(app *App) {
		app.runFunc = runFunc
	}
}

func WithSilence(silence bool) Option {
	return func(app *App) {
		app.silence = silence
	}
}

func WithUseConfig(useConfig bool) Option {
	return func(app *App) {
		app.useConfig = useConfig
	}
}

func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(app *App) {
		app.args = args
	}
}

func WithDefaultValidArgs() Option {
	return func(app *App) {
		app.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any args, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		}
	}
}

// NewApp create a new app.
func NewApp(use string, short string, opts ...Option) *App {
	app := &App{
		use:       use,
		short:     short,
		silence:   true,
		useConfig: true,
	}

	for _, opt := range opts {
		opt(app)
	}

	app.buildCmd()

	return app
}

func (a *App) buildCmd() {
	cmd := cobra.Command{
		Use:           FormatUseName(a.use),
		Short:         a.short,
		Long:          a.long,
		SilenceUsage:  a.silence,
		SilenceErrors: a.silence,
		Args:          a.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	cliflag.InitFlags(cmd.Flags())

	if len(a.commands) > 0 {
		for _, command := range a.commands {
			cmd.AddCommand(command.cobraCommand())
		}

		cmd.SetHelpCommand(helpCommand(FormatUseName(a.use)))
	}

	if a.runFunc != nil {
		cmd.RunE = a.runE
	}

	var namedFlagSets cliflag.NamedFlagSets
	if a.options != nil {
		namedFlagSets = a.options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	if a.useConfig {
		addConfigFlag(a.use, namedFlagSets.FlagSet("global"))
	}

	addCmdTemplate(&cmd, namedFlagSets)
	a.cmd = &cmd
}

func (a *App) runE(cmd *cobra.Command, args []string) error {

	printWorkingDir()
	cliflag.PrintFlags(cmd.Flags())

	// Use config file
	if a.useConfig {
		// Merge options and config file
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		if err := viper.Unmarshal(a.options); err != nil {
			return err
		}
	}

	if a.options != nil {
		if err := a.applyOptionRules(); err != nil {
			return err
		}
	}

	if a.runFunc != nil {
		return a.runFunc(a.use)
	}

	return nil
}

func (a *App) applyOptionRules() error {
	if completableOptions, ok := a.options.(CompletableOptions); ok {
		if err := completableOptions.Complete(); err != nil {
			return err
		}
	}
	// TODO 需要一个 error 包
	if errs := a.options.Validate(); len(errs) > 0 {
		return errs[0]
	}

	if printableOptions, ok := a.options.(PrintableOptions); ok && !a.silence {
		fmt.Printf("%v Config: `%s`", progressMessage, printableOptions.String())

	}

	return nil
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%s \n", color.RedString("Error: %v", err.Error()))
		os.Exit(1)
	}
}

func (a App) Command() *cobra.Command {
	return a.cmd
}

func printWorkingDir() {
	wd, _ := os.Getwd()
	fmt.Println(color.GreenString("\n%s\n", "======== WorkingDir ========"))
	fmt.Println(wd)
}
