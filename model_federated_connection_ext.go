package jossoappi

// Transforms the federated connection's channelA (FederatedChannelDTO) into a IdentityProviderChannelDTO
func (f *FederatedConnectionDTO) GetIDPChannel() (*IdentityProviderChannelDTO, error) {
	c := f.GetChannelA()
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

	idpc.SetPreferred(c.AdditionalProperties["preferred"].(bool))

	if idpc.GetOverrideProviderSetup() {
		idpc.SetSignatureHash(c.AdditionalProperties["signaturehash"].(string))
		idpc.SetMessageTtl(c.AdditionalProperties["messagettl"].(int32))
		idpc.SetMessageTtlTolerance(c.AdditionalProperties["messagettltolerance"].(int32))
		AccountLinkagePolicyDTO := toAccountLinkagePolicyDTO(c.AdditionalProperties["accountlinkagepolicy"].(map[string]interface{}))
		idpc.SetAccountLinkagePolicy(AccountLinkagePolicyDTO)
		idpc.SetEnableProxyExtension(c.AdditionalProperties["enableproxyextension"].(bool))
		IdentityMappingPolicyDTO := toIdentityMappingPolicyDTO(c.AdditionalProperties["identitymappingpolicy"].(map[string]interface{}))
		idpc.SetIdentityMappingPolicy(IdentityMappingPolicyDTO)
		idpc.SetSignAuthenticationRequests(c.AdditionalProperties["signauthenticationrequests"].(bool))
		idpc.SetWantAssertionSigned(c.AdditionalProperties["wantassertionsigned"].(bool))
	}
	return &idpc, nil
}

// Transforms the IdentityProviderChannelDTO into a FederatedChannel and sets it into channelA
func (f *FederatedConnectionDTO) SetIDPChannel(idpc *IdentityProviderChannelDTO) error {

	var c FederatedChannelDTO

	// TODO
	c.SetActiveBindings(idpc.GetActiveBindings())
	c.SetActiveProfiles(idpc.GetActiveProfiles())
	c.SetDescription(idpc.GetDescription())
	c.SetDisplayName(c.GetDisplayName())
	c.SetElementId(c.GetElementId())
	c.SetId(c.GetId())
	c.SetLocation(c.GetLocation())
	c.SetName(c.GetName())
	c.SetOverrideProviderSetup(c.GetOverrideProviderSetup())

	c.AdditionalProperties["preferred"] = idpc.GetPreferred()

	if idpc.GetOverrideProviderSetup() {
		c.AdditionalProperties["signaturehash"] = idpc.GetSignatureHash()
		// TODO.(string))
		/*
			idpc.SetMessageTtl(c.AdditionalProperties["messagettl"].(int32))
			idpc.SetMessageTtlTolerance(c.AdditionalProperties["messagettltolerance"].(int32))
			AccountLinkagePolicyDTO := toAccountLinkagePolicy(c.AdditionalProperties["accountlinkagepolicy"].(map[string]interface{}))
			idpc.SetAccountLinkagePolicy(AccountLinkagePolicyDTO)
			idpc.SetEnableProxyExtension(c.AdditionalProperties["enableproxyextension"].(bool))
			IdentityMappingPolicyDTO := toIdentityMappingPolicy(c.AdditionalProperties["identitymappingpolicy"].(map[string]interface{}))
			idpc.SetIdentityMappingPolicy(IdentityMappingPolicyDTO)
			idpc.SetSignAuthenticationRequests(c.AdditionalProperties["signauthenticationrequests"].(bool))
			idpc.SetWantAssertionSigned(c.AdditionalProperties["wantassertionsigned"].(bool))
		*/
	}

	f.SetChannelA(c)
	return nil
}

// Transforms IdentityMappingPolicyDTO a map
func toIdentityMappingPolicyMap(dto IdentityMappingPolicyDTO) *map[string]interface{} {

	props := make(map[string]interface{})

	props["custommapper"] = dto.GetCustomMapper()

	// TODO :

	return &props

}

// Transforms a map into an IdentityMappingPolicyDTO
func toIdentityMappingPolicyDTO(props map[string]interface{}) IdentityMappingPolicyDTO {
	dto := NewIdentityMappingPolicyDTO()
	dto.SetCustomMapper((props["custommapper"].(string)))
	dto.SetElementId((props["elementid"].(string)))
	dto.SetId((props["id"].(int64)))
	dto.SetMappingType((props["mappingtype"].(string)))
	dto.SetName((props["name"].(string)))
	dto.SetUseLocalId((props["uselocalid"].(bool)))
	return *dto
}

func toAccountLinkagePolicyMap(dto AccountLinkagePolicyDTO) *map[string]interface{} {
	// TODO
	return nil
}

func toAccountLinkagePolicyDTO(props map[string]interface{}) AccountLinkagePolicyDTO {
	dto := NewAccountLinkagePolicyDTO()
	dto.SetCustomLinkEmitter((props["customlinkemitter"].(string)))
	dto.SetElementId((props["elementid"].(string)))
	dto.SetId((props["id"].(int64)))
	dto.SetLinkEmitterType((props["linkemittertype"].(string)))
	dto.SetName((props["name"].(string)))
	return *dto
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

	AuthenticationAssertionEmissionPolicyDTO := toEmissionPolicy(c.AdditionalProperties["emissionpolicy"].(map[string]interface{}))
	spc.SetEmissionPolicy(AuthenticationAssertionEmissionPolicyDTO)
	spc.SetRestrictedRoles(c.AdditionalProperties["restrictedroles"].([]string))
	spc.SetRequiredRoles(c.AdditionalProperties["requiredroles"].([]string))
	spc.SetMessageTtl(c.AdditionalProperties["messageTtl"].(int32))
	spc.SetRequiredRolesMatchMode(c.AdditionalProperties["requiredrolesmatchmode"].(int32))
	spc.SetRestrictedRolesMatchMode(c.AdditionalProperties["restrictedrolesmatchmode"].(int32))
	spc.SetEncryptAssertionAlgorithm(c.AdditionalProperties["encryptassertionalgorithm"].(string))
	spc.SetIgnoreRequestedNameIDPolicy(c.AdditionalProperties["ignorerequestednameidpolicy"].(bool))
	subjectNameIdDTO := toSubjectNameIDPolicy(c.AdditionalProperties["subjectnameidpolicy"].(map[string]interface{}))
	spc.SetSubjectNameIDPolicy(subjectNameIdDTO)
	spc.SetEncryptAssertion(c.AdditionalProperties["encryptassertion"].(bool))
	spc.SetSignatureHash(c.AdditionalProperties["signaturehash"].(string))
	AttributeProfileDTO := toAttributeProfile(c.AdditionalProperties["attributeprofile"].(map[string]interface{}))
	spc.SetAttributeProfile(AttributeProfileDTO)
	AuthenticationContractDTO := toAuthenticationContract(c.AdditionalProperties["authenticationcontract"].(map[string]interface{}))
	spc.SetAuthenticationContract(AuthenticationContractDTO)
	spc.SetWantAuthnRequestsSigned(c.AdditionalProperties["wantauthnrequestssigned"].(bool))

	return &spc, nil
}

func toEmissionPolicy(props map[string]interface{}) AuthenticationAssertionEmissionPolicyDTO {
	dto := NewAuthenticationAssertionEmissionPolicyDTO()
	dto.SetElementId(props["elementid"].(string))
	dto.SetId(props["id"].(int64))
	dto.SetName(props["name"].(string))
	return *dto
}

func toAuthenticationContract(props map[string]interface{}) AuthenticationContractDTO {
	dto := NewAuthenticationContractDTO()
	dto.SetElementId(props["elementid"].(string))
	dto.SetId(props["id"].(int64))
	dto.SetName(props["name"].(string))
	return *dto
}

func toAttributeProfile(props map[string]interface{}) AttributeProfileDTO {
	dto := NewAttributeProfileDTO()
	dto.SetElementId(props["elementid"].(string))
	dto.SetId(props["id"].(int64))
	dto.SetName(props["name"].(string))
	dto.SetProfileType(props["profiletype"].(string))
	return *dto
}

func toSubjectNameIDPolicy(props map[string]interface{}) SubjectNameIdentifierPolicyDTO {
	dto := NewSubjectNameIdentifierPolicyDTO()
	dto.SetDescriptionKey(props["descriptionkey"].(string))
	dto.SetId(props["id"].(string))
	dto.SetName(props["name"].(string))
	dto.SetSubjectAttribute(props["subjectattribute"].(string))
	dto.SetType(props["type"].(string))
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
