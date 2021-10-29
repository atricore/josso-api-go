package jossoappi

func (idpCfg SamlR2IDPConfigDTO) ToProviderConfig() (*ProviderConfigDTO, error) {
	var cfg ProviderConfigDTO
	cfg.AdditionalProperties = make(map[string]interface{})
	// Build specific type
	cfg.AdditionalProperties["@id"] = idpCfg.AdditionalProperties["@id"]
	cfg.AdditionalProperties["@c"] = ".SamlR2IDPConfigDTO"

	cfg.Description = idpCfg.Description
	cfg.DisplayName = idpCfg.DisplayName
	cfg.ElementId = idpCfg.ElementId
	cfg.Name = idpCfg.Name
	cfg.AdditionalProperties["useSampleStore"] = idpCfg.UseSampleStore
	cfg.AdditionalProperties["useSystemStore"] = idpCfg.UseSystemStore

	if !*idpCfg.UseSampleStore && !*idpCfg.UseSystemStore {
		storeProps := toKeyStoreMap(idpCfg.GetSigner())
		cfg.AdditionalProperties["signer"] = storeProps
		cfg.AdditionalProperties["encrypter"] = storeProps["@id"]
	}

	return &cfg, nil

}
