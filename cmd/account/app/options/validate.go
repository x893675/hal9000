package options

func (a *AccountServiceOptions) Validate() []error {
	var errors []error
	errors = append(errors, a.DatabaseOptions.Validate()...)
	return errors
}
