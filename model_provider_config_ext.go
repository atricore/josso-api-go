package jossoappi

import (
	"errors"
	"fmt"
)

func (cfg ProviderConfigDTO) GetStoreProps() (map[string]interface{}, int, error) {
	var storeProps map[string]interface{}
	var storeId int
	var ok bool

	if storeId, ok = cfg.AdditionalProperties["encrypter"].(int); ok {
		storeProps = cfg.AdditionalProperties["signer"].(map[string]interface{})
	} else if storeId, ok = cfg.AdditionalProperties["signer"].(int); ok {
		storeProps = cfg.AdditionalProperties["encrypter"].(map[string]interface{})
	} else {
		return storeProps, storeId, fmt.Errorf("config does not have encrypter/signer ?")
	}

	return storeProps, storeId, nil
}

func (cfg ProviderConfigDTO) ToSamlR2SPConfig() (*SamlR2SPConfigDTO, error) {
	var smlr2 SamlR2SPConfigDTO
	smlr2.AdditionalProperties = make(map[string]interface{})
	// Build specific type
	smlr2.AdditionalProperties["@id"] = cfg.AdditionalProperties["@id"]
	smlr2.AdditionalProperties["@c"] = ".SamlR2SPConfigDTO"

	smlr2.Description = cfg.Description
	smlr2.ElementId = cfg.ElementId
	smlr2.Id = cfg.Id
	smlr2.Name = cfg.Name

	if !*smlr2.UseSampleStore && !*smlr2.UseSystemStore {
		storeProps := toKeyStoreMap(smlr2.GetSigner())
		smlr2.AdditionalProperties["signer"] = storeProps
		smlr2.AdditionalProperties["encrypter"] = storeProps["@id"]
	}

	return &smlr2, nil

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
