/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package flag

import (
	goflag "flag"

	log "cooool-blog-api/pkg/rollinglog"

	"github.com/spf13/pflag"
)

type NamedFlagSets struct {
	Order    []string
	FlagSets map[string]*pflag.FlagSet
}

// NewNamedFlagSets create NamedFlagSets
func NewNamedFlagSets() *NamedFlagSets {
	return &NamedFlagSets{
		Order:    make([]string, 0),
		FlagSets: make(map[string]*pflag.FlagSet),
	}
}

// FlagSet returns the flag set with the given name and adds it to the
// ordered name list if it is not in there yet.
func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if nfs.FlagSets == nil {
		nfs.FlagSets = map[string]*pflag.FlagSet{}
	}
	if _, ok := nfs.FlagSets[name]; !ok {
		nfs.FlagSets[name] = pflag.NewFlagSet(name, pflag.ExitOnError)
		nfs.Order = append(nfs.Order, name)
	}
	return nfs.FlagSets[name]
}

// InitFlags normalizes, parses, then logs the command line flags.
func InitFlags(fs *pflag.FlagSet) {
	fs.SetNormalizeFunc(WordReplaceNormalizeFunc)
	fs.AddGoFlagSet(goflag.CommandLine)
}

// PrintFlags logs the flags in the flagSet.
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Debug("Flag value has been parsed")
	})
}
