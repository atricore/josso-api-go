package jossoappi

func (spCfg SamlR2SPConfigDTO) ToProviderConfig() (*ProviderConfigDTO, error) {
	var cfg ProviderConfigDTO
	cfg.AdditionalProperties = make(map[string]interface{})
	// Build specific type
	cfg.AdditionalProperties["@id"] = spCfg.AdditionalProperties["@id"]
	cfg.AdditionalProperties["@c"] = ".SamlR2SPConfigDTO"

	cfg.Description = spCfg.Description
	cfg.DisplayName = spCfg.DisplayName
	cfg.ElementId = spCfg.ElementId
	cfg.Name = spCfg.Name
	cfg.AdditionalProperties["useSampleStore"] = spCfg.UseSampleStore
	cfg.AdditionalProperties["useSystemStore"] = spCfg.UseSystemStore

	if !*spCfg.UseSampleStore && !*spCfg.UseSystemStore {
		storeProps := toKeyStoreMap(spCfg.GetSigner())
		cfg.AdditionalProperties["signer"] = storeProps
		cfg.AdditionalProperties["encrypter"] = storeProps["@id"]
	}

	return &cfg, nil

}
