package options

func (s *ServerRunOptions) Validate() []error {
	var errors []error
	errors = append(errors, s.Validate()...)
	return errors
}
