package app

import (
	"blog-go/pkg/cli/flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	progressMessage = color.GreenString("==>")

	commandDesc = `Welcome to api-server`

	usageTemplate = fmt.Sprintf(`%s{{if .Runnable}}
  %s{{end}}{{if .HasAvailableSubCommands}}
  %s{{end}}{{if gt (len .Aliases) 0}}

%s
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

%s
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

%s{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  %s {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

%s
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

%s
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

%s{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "%s --help" for more information about a command.{{end}}
`,
		color.CyanString("Usage:"),
		color.GreenString("{{.UseLine}}"),
		color.GreenString("{{.CommandPath}} [command]"),
		color.CyanString("Aliases:"),
		color.CyanString("Examples:"),
		color.CyanString("Available Commands:"),
		color.GreenString("{{rpad .Name .NamePadding }}"),
		color.CyanString("Flags:"),
		color.CyanString("Global Flags:"),
		color.CyanString("Additional help topics:"),
		color.GreenString("{{.CommandPath}} [command]"),
	)
)

type App struct {
	// 应用名称
	use   string
	short string
	long  string
	// flags
	flags CliFlags
	cmd   *cobra.Command
	// 子命令
	commands []*Command
	// 非标志参数验证函数
	args cobra.PositionalArgs
	// 允许的非标志参数
	validArgs     []string
	runFunc       RunFunc
	silenceUsage  bool
	silenceErrors bool
}

type Option func(*App)

type RunFunc func(basename string) error

func WithLong(desc string) Option {
	return func(app *App) {
		app.long = desc
	}
}

func WithFlags(opts CliFlags) Option {
	return func(app *App) {
		app.flags = opts
	}
}

func WithArgs(args cobra.PositionalArgs) Option {
	return func(app *App) {
		app.args = args
	}
}

func WithValidArgs(validArgs []string) Option {
	return func(app *App) {
		app.validArgs = validArgs
	}
}

func WithRunFunc(runFunc RunFunc) Option {
	return func(app *App) {
		app.runFunc = runFunc
	}
}

func WithSilenceUsage(silenceUsage bool) Option {
	return func(app *App) {
		app.silenceUsage = silenceUsage
	}
}

func WithSilenceErrors(silenceErrors bool) Option {
	return func(app *App) {
		app.silenceUsage = silenceErrors
	}
}

func NewApp(use string, short string, opts ...Option) *App {
	app := &App{
		use:           use,
		short:         short,
		long:          commandDesc,
		silenceUsage:  true,
		silenceErrors: true,
	}

	for _, opt := range opts {
		opt(app)
	}

	app.buildCmd()

	return app
}

func (a *App) buildCmd() {
	cmd := cobra.Command{
		Use:           FormatBaseName(a.use),
		Short:         a.short,
		Long:          a.long,
		SilenceUsage:  a.silenceUsage,
		SilenceErrors: a.silenceErrors,
		Args:          a.args,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	flag.InitFlags(cmd.Flags())
	cmd.Flags().SortFlags = true

	// 如果子命令不为空，则追加子命令
	if len(a.commands) > 0 {
		for _, command := range a.commands {
			cmd.AddCommand(command.cobraCommand())
		}
	}

	var namedFlagSets flag.FlagSets
	if a.flags != nil {
		namedFlagSets = a.flags.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	cmd.RunE = a.runE
	a.cmd = &cmd
}

func (a *App) runE(cmd *cobra.Command, args []string) error {
	fmt.Println(color.GreenString("api server is start"))
	fmt.Println(args)
	if a.runFunc != nil {
		return a.runFunc(a.use)
	}

	return nil
}

// GetCmd 返回 app 的 cmd
func (a *App) GetCmd() *cobra.Command {
	return a.cmd
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		log.Fatalln(color.RedString(err.Error()))
	}
}
