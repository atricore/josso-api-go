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
	spCfg := NewSamlR2SPConfigDTO()
	spCfg.AdditionalProperties = make(map[string]interface{})

	// @id and @c properties
	class := cfg.AdditionalProperties["@c"]
	if class == nil {
		return spCfg, errors.New("class property not found (@c)")
	}

	if class != ".SamlR2SPConfigDTO" {
		return spCfg, fmt.Errorf("invalid class %s", class)
	}

	// Build specific type
	spCfg.AdditionalProperties["@id"] = cfg.AdditionalProperties["@id"]
	spCfg.AdditionalProperties["@c"] = class

	spCfg.SetDescription(cfg.GetDescription())
	spCfg.SetDisplayName(cfg.GetDisplayName())
	spCfg.SetElementId(cfg.GetElementId())
	spCfg.SetId(cfg.GetId())
	spCfg.SetName(cfg.GetName())
	spCfg.SetUseSampleStore(AsBool(cfg.AdditionalProperties["useSampleStore"], false))
	spCfg.SetUseSystemStore(AsBool(cfg.AdditionalProperties["useSystemStore"], false))

	if !*spCfg.UseSampleStore && !*spCfg.UseSystemStore {
		// Get signer/encrypter
		storeProps, storeId, err := cfg.GetStoreProps()
		if err != nil {
			return spCfg, err
		}

		if storeProps["@id"].(int) != storeId {
			return spCfg, fmt.Errorf("inconsistent config Ids %d, %d", storeId, storeProps["@id"].(int))
		}

		store := toKeyStoreDTO(storeId, storeProps)
		spCfg.Signer = store
		spCfg.Encrypter = store

	}

	return spCfg, nil

}

func (cfg ProviderConfigDTO) ToSamlR2IDPConfig() (*SamlR2IDPConfigDTO, error) {

	idpCfg := NewSamlR2IDPConfigDTO()
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
