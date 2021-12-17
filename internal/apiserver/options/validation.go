package options

func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.APISServerOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)

	return errs
}
