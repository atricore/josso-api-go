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

func (f *FederatedConnectionDTO) GetSPChannel() (*InternalSaml2ServiceProviderChannelDTO, error) {
	c := f.GetChannelA()
	var spc InternalSaml2ServiceProviderChannelDTO

	spc.SetId(c.GetId())
	/// ....

	// requiredRoles
	spc.SetMessageTtl(c.AdditionalProperties["messageTtl"].(int32))

	//idpFc.SetEnableProxyExtension()

	return &spc, nil
}

func (f *FederatedConnectionDTO) SetSPChannel(idpChannel *InternalSaml2ServiceProviderChannelDTO) error {
	return nil
}
