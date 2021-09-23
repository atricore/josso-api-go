package jossoappi

import (
	"errors"
	"fmt"
)

func (idpCfg *SamlR2IDPConfigDTO) ToProviderConfig() (*ProviderConfigDTO, error) {
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

func (cfg ProviderConfigDTO) ToSamlR2IDPConfig() (*SamlR2IDPConfigDTO, error) {

	var idpCfg *SamlR2IDPConfigDTO

	idpCfg.AdditionalProperties = make(map[string]interface{})

	// @id and @c properties
	class := cfg.AdditionalProperties["@c"]
	if class == nil {
		return idpCfg, errors.New("class property not found (@c)")
	}

	if class != ".SamlR2IDPConfigDTO" {
		return idpCfg, fmt.Errorf("invalid class %s", class)
	}

	// Build specific type
	idpCfg.AdditionalProperties["@id"] = cfg.AdditionalProperties["@id"]
	idpCfg.AdditionalProperties["@c"] = class

	idpCfg.Description = cfg.Description
	idpCfg.DisplayName = cfg.DisplayName
	idpCfg.ElementId = cfg.ElementId
	idpCfg.Name = cfg.Name
	idpCfg.UseSampleStore = PtrBool(cfg.AdditionalProperties["useSampleStore"].(bool))
	idpCfg.UseSystemStore = PtrBool(cfg.AdditionalProperties["useSystemStore"].(bool))

	if !*idpCfg.UseSampleStore && !*idpCfg.UseSystemStore {
		// Get signer/encrypter
		storeProps, storeId, err := cfg.GetStoreProps()
		if err != nil {
			return idpCfg, err
		}

		if storeProps["@id"].(int) != storeId {
			return idpCfg, fmt.Errorf("inconsistent config Ids %d, %d", storeId, storeProps["@id"].(int))
		}

		store := toKeyStoreDTO(storeId, storeProps)
		idpCfg.Signer = store
		idpCfg.Encrypter = store

	}

	return idpCfg, nil
}
