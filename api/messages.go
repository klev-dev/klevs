package api

func (r *PublishRequest) Validate() error {
	return r.Log.Validate()
}

func (r *ConsumeRequest) Validate() error {
	return r.Log.Validate()
}
