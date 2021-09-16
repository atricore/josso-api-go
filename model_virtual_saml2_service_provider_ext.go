package jossoappi

func (p *VirtualSaml2ServiceProviderDTO) GetSamlR2IDPConfig() (*SamlR2IDPConfigDTO, error) {
	return toSamlR2IDPConfig(p.GetConfig())
}

func (p *VirtualSaml2ServiceProviderDTO) SetSamlR2IDPConfig(idpCfg *SamlR2IDPConfigDTO) error {
	cfg, err := toProviderConfig(*idpCfg)

	if err != nil {
		return err
	}
	p.SetConfig(*cfg)
	return nil

}
