package options

import (
	genericoptions "blog-go/internal/pkg/options"
	"blog-go/pkg/cli/flag"
	"encoding/json"
)

type Options struct {
	APISServerOptions *genericoptions.APIServerOptions
	MySQLOptions      *genericoptions.MySQLOptions
}

func NewOptions() *Options {
	return &Options{
		APISServerOptions: genericoptions.NewServerOptions(),
		MySQLOptions:      genericoptions.NewMySQLOptions(),
	}
}

func (o *Options) Flags() (fss flag.FlagSets) {
	o.APISServerOptions.AddFlags(fss.FlagSet("api-server"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))

	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
