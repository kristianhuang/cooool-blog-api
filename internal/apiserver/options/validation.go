/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

// Validate checks options and return errors.
func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.InsecureServingOptions.Validate()...)
	errs = append(errs, o.RedisOptions.Validate()...)

	return errs
}
