/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package app

import "blog-go/pkg/cli/flag"

type CliFlags interface {
	Flags() (fss flag.FlagSets)
	Validate() []error
}

type Config interface {
}

// ConfigurableOptions abstracts configuration flags for reading parameters
// from a configuration file.
type ConfigurableOptions interface {
	// ApplyFlags parsing parameters from the command line or configuration file
	// to the flags instance.
	ApplyFlags() []error
}

// CompletableOptions abstracts flags, which can be completed.
type CompletableOptions interface {
	Complete() error
}

// PrintableOptions abstracts flags, which can be printed.
type PrintableOptions interface {
	String() string
}
