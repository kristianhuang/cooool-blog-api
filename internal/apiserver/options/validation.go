/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.APISServerOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)

	return errs
}
