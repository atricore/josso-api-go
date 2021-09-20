package jossoappi

// Ignore ConnectionA and ConnectionB properties

func (f *FederatedConnectionDTO) GetIDPChannel() (*IdentityProviderChannelDTO, error) {

	c := f.GetChannelB()
	var idpc IdentityProviderChannelDTO

	idpc.SetId(c.GetId())

	//idpFc.SetEnableProxyExtension()

	return &idpc, nil
}

func (f *FederatedConnectionDTO) SetIDPChannel(idpChannel *IdentityProviderChannelDTO) error {
	return nil
}

// IDP Side, has an SP channel
func (f *FederatedConnectionDTO) GetSPChannel() (*InternalSaml2ServiceProviderChannelDTO, error) {
	c := f.GetChannelA()
	var spc InternalSaml2ServiceProviderChannelDTO

	spc.SetId(c.GetId())
	/// ....

	// requiredRoles
	spc.SetMessageTtl(c.AdditionalProperties["messageTtl"].(int32))

	// Mapping array
	spc.SetActiveBindings(c.AdditionalProperties["activeBindings"].([]string))

	// Mapping struct / complext type
	attributeProfileDTO := toAttributeProfile(c.AdditionalProperties["attributeProfile"].(map[string]interface{}))
	spc.SetAttributeProfile(attributeProfileDTO)

	// Mapping struct / complex type
	subjectNameIdDTO := toSubjectNameIDPolicy(c.AdditionalProperties["subjectNameIDPolicy"].(map[string]interface{}))
	spc.SetSubjectNameIDPolicy(subjectNameIdDTO)

	//idpFc.SetEnableProxyExtension()

	return &spc, nil
}

func toAttributeProfile(props map[string]interface{}) AttributeProfileDTO {
	dto := NewAttributeProfileDTO()

	dto.SetName(props["Name"].(string))
	// TODO :
	return *dto

}

func toSubjectNameIDPolicy(props map[string]interface{}) SubjectNameIdentifierPolicyDTO {
	dto := NewSubjectNameIdentifierPolicyDTO()
	dto.SetType(props["type"].(string))
	dto.SetName(props["name"].(string))
	dto.SetSubjectAttribute(props["subjectAttribute"].(string))
	dto.SetDescriptionKey(props["descriptionKey"].(string))
	return *dto
}

func (f *FederatedConnectionDTO) SetSPChannel(idpChannel *InternalSaml2ServiceProviderChannelDTO) error {
	return nil
}

func addFederatedConnection(fcs []FederatedConnectionDTO,
	target string,
	spChannel *InternalSaml2ServiceProviderChannelDTO,
	idpChannel *IdentityProviderChannelDTO) ([]FederatedConnectionDTO, error) {

	// Create new Federated Connection
	var fc FederatedConnectionDTO
	fc.SetName(target)
	fc.SetIDPChannel(idpChannel)
	fc.SetSPChannel(spChannel)

	fcs = append(fcs, fc)

	return fcs, nil
}

func removeFederatedConnection(fcs []FederatedConnectionDTO, target string) ([]FederatedConnectionDTO, error) {

	var newFcs []FederatedConnectionDTO
	for _, fc := range fcs {
		if fc.GetName() != target {
			newFcs = append(newFcs, fc)
		}
	}

	return newFcs, nil
}
