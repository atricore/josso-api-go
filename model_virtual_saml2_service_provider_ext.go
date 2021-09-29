package jossoappi

func (p *VirtualSaml2ServiceProviderDTO) GetSamlR2IDPConfig() (*SamlR2IDPConfigDTO, error) {
	return p.GetConfig().ToSamlR2IDPConfig()
}

func (p *VirtualSaml2ServiceProviderDTO) SetSamlR2IDPConfig(idpCfg *SamlR2IDPConfigDTO) error {
	cfg, err := idpCfg.ToProviderConfig()

	if err != nil {
		return err
	}
	p.SetConfig(*cfg)
	return nil

}

// TODO : Add/Remove federated connection (see External SAML 2 SP)
