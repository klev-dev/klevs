package api

func (r *LogsListRequest) Validate() error {
	return nil
}

func (r *LogsCreateRequest) Validate() error {
	return ValidateName(r.Name)
}

func (r *LogsDeleteRequest) Validate() error {
	return ValidateName(r.Name)
}
