package options


func (a *AuthServiceOptions) Validate() []error {
	var errors []error
	errors = append(errors, a.DatabaseOptions.Validate()...)
	return errors
}