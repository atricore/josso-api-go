package jossoappi

func (idpCfg SamlR2SPConfigDTO) ToProviderConfig() (*ProviderConfigDTO, error) {
	var prc ProviderConfigDTO
	prc.AdditionalProperties = make(map[string]interface{})
	// Build specific type
	prc.AdditionalProperties["@id"] = idpCfg.AdditionalProperties["@id"]
	prc.AdditionalProperties["@c"] = ".SamlR2IDPConfigDTO"

	prc.Description = idpCfg.Description
	prc.DisplayName = idpCfg.DisplayName
	prc.ElementId = idpCfg.ElementId
	prc.Name = idpCfg.Name
	prc.AdditionalProperties["useSampleStore"] = idpCfg.UseSampleStore
	prc.AdditionalProperties["useSystemStore"] = idpCfg.UseSystemStore

	if !*idpCfg.UseSampleStore && !*idpCfg.UseSystemStore {
		storeProps := toKeyStoreMap(idpCfg.GetSigner())
		prc.AdditionalProperties["signer"] = storeProps
		prc.AdditionalProperties["encrypter"] = storeProps["@id"]
	}

	return &prc, nil

}
