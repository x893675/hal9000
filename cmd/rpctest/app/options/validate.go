package options


func (s *ServiceRunOptions) Validate() []error {
	var errors []error
	errors = append(errors, s.MySQLOptions.Validate()...)
	return errors
}