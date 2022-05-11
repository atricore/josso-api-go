package jossoappi

// Transforms the federated connection's channelA (FederatedChannelDTO) into a IdentityProviderChannelDTO
func (f *FederatedConnectionDTO) GetIDPChannel() (*IdentityProviderChannelDTO, error) {
	c := f.GetChannelB()
	var idpc IdentityProviderChannelDTO

	idpc.SetActiveBindings(c.GetActiveBindings())
	idpc.SetActiveProfiles(c.GetActiveProfiles())
	idpc.SetDescription(c.GetDescription())
	idpc.SetDisplayName(c.GetDisplayName())
	idpc.SetElementId(c.GetElementId())
	idpc.SetId(c.GetId())
	idpc.SetLocation(c.GetLocation())
	idpc.SetName(c.GetName())
	idpc.SetOverrideProviderSetup(c.GetOverrideProviderSetup())

	prefered := false
	if c.AdditionalProperties["preferred"] != nil {
		prefered = (c.AdditionalProperties["preferred"].(bool))
	}
	idpc.SetPreferred(prefered)

	if idpc.GetOverrideProviderSetup() {
		idpc.SetSignatureHash(c.AdditionalProperties["signatureHash"].(string))
		idpc.SetMessageTtl(c.AdditionalProperties["messageTtl"].(int32))
		idpc.SetMessageTtlTolerance(c.AdditionalProperties["messageTtlTolerance"].(int32))
		accountLinkage := toAccountLinkagePolicyDTO(c.AdditionalProperties["accountLinkagePolicy"].(map[string]interface{}))
		idpc.SetAccountLinkagePolicy(accountLinkage)
		idpc.SetEnableProxyExtension(c.AdditionalProperties["enableProxyExtension"].(bool))
		idMapping := toIdentityMappingPolicyDTO(c.AdditionalProperties["identityMappingPolicy"].(map[string]interface{}))
		idpc.SetIdentityMappingPolicy(idMapping)
		idpc.SetSignAuthenticationRequests(c.AdditionalProperties["signAuthenticationRequests"].(bool))
		idpc.SetWantAssertionSigned(c.AdditionalProperties["wantAssertionSigned"].(bool))
	}
	return &idpc, nil
}

// Transforms the IdentityProviderChannelDTO into a FederatedChannel and sets it into channelA
func (f *FederatedConnectionDTO) SetIDPChannel(idpc *IdentityProviderChannelDTO) error {

	var c FederatedChannelDTO

	c.SetActiveBindings(idpc.GetActiveBindings())
	c.SetActiveProfiles(idpc.GetActiveProfiles())
	c.SetDescription(idpc.GetDescription())
	c.SetDisplayName(idpc.GetDisplayName())
	c.SetElementId(idpc.GetElementId())
	c.SetId(idpc.GetId())
	c.SetLocation(idpc.GetLocation())
	c.SetName(idpc.GetName())
	c.SetOverrideProviderSetup(idpc.GetOverrideProviderSetup())

	c.AdditionalProperties = make(map[string]interface{})
	c.AdditionalProperties["preferred"] = idpc.GetPreferred()
	c.AdditionalProperties["@c"] = ".IdentityProviderChannelDTO"

	if idpc.GetOverrideProviderSetup() {
		c.AdditionalProperties["signatureHash"] = idpc.GetSignatureHash()
		c.AdditionalProperties["messageTtl"] = idpc.GetMessageTtl()
		c.AdditionalProperties["messageTtlTolerance"] = idpc.GetMessageTtlTolerance()
		c.AdditionalProperties["accountLinkagePolicy"] = toAccountLinkagePolicyMap(idpc.GetAccountLinkagePolicy())
		c.AdditionalProperties["enableProxyExtension"] = idpc.GetEnableProxyExtension()
		c.AdditionalProperties["identityMappingPolicy"] = toIdentityMappingPolicyMap(idpc.GetIdentityMappingPolicy())
		c.AdditionalProperties["signAuthenticationRequests"] = idpc.GetSignAuthenticationRequests()
		c.AdditionalProperties["wantAssertionSigned"] = idpc.GetWantAssertionSigned()
	}

	f.SetChannelB(c)
	return nil
}

// Transforms IdentityMappingPolicyDTO a map
func toIdentityMappingPolicyMap(dto IdentityMappingPolicyDTO) *map[string]interface{} {

	props := make(map[string]interface{})

	props["customMapper"] = dto.GetCustomMapper()
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["mappingType"] = dto.GetMappingType()
	props["name"] = dto.GetName()
	props["useLocalId"] = dto.GetUseLocalId()
	return &props
}

// Transforms a map into an IdentityMappingPolicyDTO
func toIdentityMappingPolicyDTO(props map[string]interface{}) IdentityMappingPolicyDTO {
	dto := NewIdentityMappingPolicyDTO()
	dto.SetCustomMapper((props["customMapper"].(string)))
	dto.SetElementId((props["elementId"].(string)))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetMappingType((props["mappingType"].(string)))
	dto.SetName((props["name"].(string)))
	dto.SetUseLocalId((props["useLocalId"].(bool)))
	return *dto
}

// Transforms AccountLinkagePolicy a map
func toAccountLinkagePolicyMap(dto AccountLinkagePolicyDTO) *map[string]interface{} {
	props := make(map[string]interface{})

	props["customLinkEmitter"] = dto.GetCustomLinkEmitter()
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["linkEmitterType"] = dto.GetLinkEmitterType()
	props["name"] = dto.GetName()
	return &props
}

// Transforms a map into an AccountLinkagePolicyDTO
func toAccountLinkagePolicyDTO(props map[string]interface{}) AccountLinkagePolicyDTO {
	dto := NewAccountLinkagePolicyDTO()
	dto.SetCustomLinkEmitter((props["customLinkEmitter"].(string)))
	dto.SetElementId((props["elementId"].(string)))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetLinkEmitterType((props["linkEmitterType"].(string)))
	dto.SetName((props["name"].(string)))
	return *dto
}

// Transforms a map into an EmissionPolicyDTO
func toEmissionPolicyDTO(props map[string]interface{}) AuthenticationAssertionEmissionPolicyDTO {
	dto := NewAuthenticationAssertionEmissionPolicyDTO()
	dto.SetElementId(props["elementId"].(string))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetName(props["name"].(string))
	return *dto
}

// Transforms a map into an AuthenticationContractDTO
func toAuthenticationContractDTO(props map[string]interface{}) AuthenticationContractDTO {
	dto := NewAuthenticationContractDTO()
	dto.SetElementId(props["elementId"].(string))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetName(props["name"].(string))
	return *dto
}

// Transforms a map into an AttributeProfileDTO
func toAttributeProfileDTO(props map[string]interface{}) AttributeProfileDTO {
	dto := NewAttributeProfileDTO()
	dto.SetElementId(props["elementId"].(string))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetName(props["name"].(string))
	dto.SetProfileType(props["profileType"].(string))
	return *dto
}

// Transforms a map into an SubjectNameIDPolicyDTO
func toSubjectNameIDPolicyDTO(props map[string]interface{}) SubjectNameIdentifierPolicyDTO {
	dto := NewSubjectNameIdentifierPolicyDTO()
	dto.SetDescriptionKey(props["descriptionKey"].(string))
	dto.SetId(props["id"].(string))
	dto.SetName(props["name"].(string))
	dto.SetSubjectAttribute(props["subjectAttribute"].(string))
	dto.SetType(props["type"].(string))
	return *dto
}

// Transforms the InternalSaml2ServiceProviderChannelDTO into a FederatedChannel and sets it into channelA
func (f *FederatedConnectionDTO) SetSPChannel(spc *InternalSaml2ServiceProviderChannelDTO) error {

	var c FederatedChannelDTO

	c.SetId(spc.GetId())
	c.SetActiveBindings(spc.GetActiveBindings())
	c.SetActiveProfiles(spc.GetActiveProfiles())
	c.SetDescription(spc.GetDescription())
	c.SetDisplayName(spc.GetDisplayName())
	c.SetElementId(spc.GetElementId())
	c.SetLocation(spc.GetLocation())
	c.SetName(spc.GetName())
	c.SetOverrideProviderSetup(spc.GetOverrideProviderSetup())

	c.AdditionalProperties = make(map[string]interface{})
	c.AdditionalProperties["@c"] = ".InternalSaml2ServiceProviderChannelDTO"

	if spc.GetOverrideProviderSetup() {
		c.AdditionalProperties["emissionPolicy"] = toEmissionPolicyMap(spc.GetEmissionPolicy())
		c.AdditionalProperties["restrictedRoles"] = spc.GetRestrictedRoles()
		c.AdditionalProperties["requiredRoles"] = spc.GetRequiredRoles()
		c.AdditionalProperties["messageTtl"] = spc.GetMessageTtl()
		c.AdditionalProperties["requiredRolesMatchMode"] = spc.GetRequiredRolesMatchMode()
		c.AdditionalProperties["restrictedRolesMatchMode"] = spc.GetRestrictedRolesMatchMode()
		c.AdditionalProperties["encryptAssertionAlgorithm"] = spc.GetEncryptAssertionAlgorithm()
		c.AdditionalProperties["ignoreRequestedNameIDPolicy"] = spc.GetIgnoreRequestedNameIDPolicy()
		c.AdditionalProperties["subjectNameIDPolicy"] = toSubjectNameIDPolicyMap(spc.GetSubjectNameIDPolicy())
		c.AdditionalProperties["encryptAssertion"] = spc.GetEncryptAssertion()
		c.AdditionalProperties["signatureHash"] = spc.GetSignatureHash()
		c.AdditionalProperties["attributeProfile"] = toAttributeProfilemap(spc.GetAttributeProfile())
		c.AdditionalProperties["authenticationContract"] = toAuthenticationContractmap(spc.GetAuthenticationContract())
		c.AdditionalProperties["wantAuthnRequestsSigned"] = spc.GetWantAuthnRequestsSigned()
	}

	f.SetChannelA(c)
	return nil

}

// IDP Side, has an SP channel
func (f *FederatedConnectionDTO) GetSPChannel() (*InternalSaml2ServiceProviderChannelDTO, error) {
	c := f.GetChannelA()
	var spc InternalSaml2ServiceProviderChannelDTO

	spc.SetId(c.GetId())
	spc.SetActiveBindings(c.GetActiveBindings())
	spc.SetActiveProfiles(c.GetActiveProfiles())
	spc.SetDescription(c.GetDescription())
	spc.SetDisplayName(c.GetDisplayName())
	spc.SetElementId(c.GetElementId())
	spc.SetLocation(c.GetLocation())
	spc.SetName(c.GetName())
	spc.SetOverrideProviderSetup(c.GetOverrideProviderSetup())

	if c.GetOverrideProviderSetup() {
		emissionPolicy := toEmissionPolicyDTO(c.AdditionalProperties["emissionPolicy"].(map[string]interface{}))
		spc.SetEmissionPolicy(emissionPolicy)
		spc.SetRestrictedRoles(c.AdditionalProperties["restrictedRoles"].([]string))
		spc.SetRequiredRoles(c.AdditionalProperties["requiredRoles"].([]string))
		spc.SetMessageTtl(c.AdditionalProperties["messageTtl"].(int32))
		spc.SetRequiredRolesMatchMode(c.AdditionalProperties["requiredRolesMatchMode"].(int32))
		spc.SetRestrictedRolesMatchMode(c.AdditionalProperties["restrictedRolesMatchMode"].(int32))
		spc.SetEncryptAssertionAlgorithm(c.AdditionalProperties["encryptAssertionAlgorithm"].(string))
		spc.SetIgnoreRequestedNameIDPolicy(c.AdditionalProperties["ignoreRequestedNameIDPolicy"].(bool))
		subjectNameId := toSubjectNameIDPolicyDTO(c.AdditionalProperties["subjectNameIDPolicy"].(map[string]interface{}))
		spc.SetSubjectNameIDPolicy(subjectNameId)
		spc.SetEncryptAssertion(c.AdditionalProperties["encryptAssertion"].(bool))
		spc.SetSignatureHash(c.AdditionalProperties["signatureHash"].(string))
		attrProfile := toAttributeProfileDTO(c.AdditionalProperties["attributeProfile"].(map[string]interface{}))
		spc.SetAttributeProfile(attrProfile)
		authnContract := toAuthenticationContractDTO(c.AdditionalProperties["authenticationContract"].(map[string]interface{}))
		spc.SetAuthenticationContract(authnContract)
		spc.SetWantAuthnRequestsSigned(c.AdditionalProperties["wantAuthnRequestsSigned"].(bool))
	}

	return &spc, nil
}

// Transforms AuthenticationContract a map
func toAuthenticationContractmap(dto AuthenticationContractDTO) *map[string]interface{} {
	props := make(map[string]interface{})
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["name"] = dto.GetName()
	return &props
}

// Transforms AttributeProfile a map
func toAttributeProfilemap(dto AttributeProfileDTO) *map[string]interface{} {
	props := make(map[string]interface{})
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["name"] = dto.GetName()
	props["profileType"] = dto.GetProfileType()
	return &props
}

// Transforms EmissionPolicy a map
func toEmissionPolicyMap(dto AuthenticationAssertionEmissionPolicyDTO) *map[string]interface{} {
	props := make(map[string]interface{})
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["name"] = dto.GetName()
	return &props
}

// Transforms SubjectNameIDPolicy a map
func toSubjectNameIDPolicyMap(dto SubjectNameIdentifierPolicyDTO) *map[string]interface{} {
	props := make(map[string]interface{})
	props["descriptionKey"] = dto.GetDescriptionKey()
	props["id"] = dto.GetId()
	props["name"] = dto.GetName()
	props["subjectAttribute"] = dto.GetSubjectAttribute()
	props["type"] = dto.GetType()
	return &props
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
	fc.AdditionalProperties = make(map[string]interface{})
	fc.AdditionalProperties["@c"] = ".FederatedConnectionDTO"

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
