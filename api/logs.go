package api

func (r *LogsListRequest) Validate() error {
	return nil
}

func (r *LogsCreateRequest) Validate() error {
	return r.Log.Validate()
}

func (r *LogsDeleteRequest) Validate() error {
	return r.Log.Validate()
}
