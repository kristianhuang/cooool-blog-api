/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"blog-api/pkg/util/path/dir"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const confFlagName = "config"

var (
	confFile string
)

func init() {
	pflag.StringVarP(&confFile, confFlagName, "c", confFile, "App config support JSON, YAML.")
}

func addConfigFlag(use string, fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(confFlagName))

	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(use), "-", "_", -1))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	cobra.OnInitialize(func() {
		if confFile != "" {
			viper.SetConfigFile(confFile)
		} else {
			viper.AddConfigPath("./config")
			viper.SetConfigType("yaml")
			if names := strings.Split(use, "-"); len(names) > 1 {
				viper.AddConfigPath(filepath.Join(dir.HomeDir(), "."+names[0]))
				viper.AddConfigPath(filepath.Join("/etc", names[0]))
			}
			viper.SetConfigName(use)
		}

		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", confFile, err)
			os.Exit(1)
		}
	})
}

func printConfig() {
	keys := viper.AllKeys()
	if len(keys) > 0 {
		fmt.Printf("%v Config items:\n", progressMessage)
		table := uitable.New()
		table.Separator = " "
		table.MaxColWidth = 80
		table.RightAlign(0)

		for _, key := range keys {
			table.AddRow(fmt.Sprintf("%s:", key), viper.Get(key))
		}

		fmt.Printf("%v", table)
	}
}
