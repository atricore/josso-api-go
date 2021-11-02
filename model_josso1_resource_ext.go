package jossoappi

func (rs *JOSSO1ResourceDTO) NewActivation(name string) (ActivationDTO, error) {
	var ac ActivationDTO
	ac.AdditionalProperties = make(map[string]interface{})
	ac.AdditionalProperties["@c"] = ".ActivationDTO"
	ac.SetName(name)
	rs.SetActivation(ac)

	return ac, nil
}

func (rs *JOSSO1ResourceDTO) RemoveActivation() bool {
	// Remove an element from : p.IdentityLookups

	if rs.Activation == nil {
		return false
	}

	rs.Activation = nil

	return true
}
