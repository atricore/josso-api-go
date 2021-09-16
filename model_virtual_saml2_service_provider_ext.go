package jossoappi

func (p *VirtualSaml2ServiceProviderDTO) GetSamlR2IDPConfig() (*SamlR2IDPConfigDTO, error) {
	return getSamlR2IDPConfig(p.GetConfig())
}
