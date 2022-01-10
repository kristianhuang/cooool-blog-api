/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package app

import "blog-api/pkg/cli/flag"

type CliOptions interface {
	Flags() (fss flag.NamedFlagSets)
	Validate() []error
}

// ConfigurableOptions abstracts configuration options for reading parameters
// from a configuration file.
type ConfigurableOptions interface {
	// ApplyFlags parsing parameters from the command line or configuration file
	// to the options instance.
	ApplyFlags() []error
}

// CompletableOptions abstracts options, which can be completed.
type CompletableOptions interface {
	Complete() error
}

// PrintableOptions abstracts options, which can be printed.
type PrintableOptions interface {
	String() string
}
