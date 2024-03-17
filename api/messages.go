package api

func (r *PublishRequest) Validate() error {
	return ValidateName(r.Name)
}

func (r *ConsumeRequest) Validate() error {
	return ValidateName(r.Name)
}
