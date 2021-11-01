package jossoappi

import (
	"fmt"
)

func (i *InternalSaml2ServiceProviderDTO) GetSamlR2SPConfig() (*SamlR2SPConfigDTO, error) {

	cfg := i.GetConfig()
	var spCfg *SamlR2SPConfigDTO
	spCfg.SetDescription(cfg.GetDescription())
	spCfg.SetDisplayName(cfg.GetDisplayName())
	spCfg.SetElementId(cfg.GetElementId())
	spCfg.SetId(cfg.GetId())
	spCfg.SetName(cfg.GetName())
	spCfg.SetUseSampleStore(cfg.AdditionalProperties["useSampleStore"].(bool))
	spCfg.SetUseSystemStore(cfg.AdditionalProperties["useSystemStore"].(bool))

	if !*spCfg.UseSampleStore && !*spCfg.UseSystemStore {
		// Get signer/encrypter
		var storeProps map[string]interface{}
		var storeId int
		var ok bool

		if storeId, ok = cfg.AdditionalProperties["encrypter"].(int); ok {
			storeProps = cfg.AdditionalProperties["signer"].(map[string]interface{})
		} else if storeId, ok = cfg.AdditionalProperties["signer"].(int); ok {
			storeProps = cfg.AdditionalProperties["encrypter"].(map[string]interface{})
		} else {
			return spCfg, fmt.Errorf("config does not have encrypter/signer ?")
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

func (i *InternalSaml2ServiceProviderDTO) SetSamlR2SPConfig(spCfg *SamlR2SPConfigDTO) error {

	var cfg ProviderConfigDTO
	cfg.AdditionalProperties = make(map[string]interface{})
	// Build specific type
	cfg.AdditionalProperties["@id"] = spCfg.AdditionalProperties["@id"]
	cfg.AdditionalProperties["@c"] = ".SamlR2IDPConfigDTO"

	cfg.Description = spCfg.Description
	cfg.DisplayName = spCfg.DisplayName
	cfg.ElementId = spCfg.ElementId
	cfg.Name = spCfg.Name
	cfg.AdditionalProperties["useSampleStore"] = spCfg.UseSampleStore
	cfg.AdditionalProperties["useSystemStore"] = spCfg.UseSystemStore

	if !*spCfg.UseSampleStore && !*spCfg.UseSystemStore {
		storeProps := toKeyStoreMap(spCfg.GetSigner()) // Assuming that both signer and encrypter have the same store
		cfg.AdditionalProperties["encrypter"] = storeProps["@id"]
		cfg.AdditionalProperties["signer"] = storeProps
	}

	return nil

}

func (p *InternalSaml2ServiceProviderDTO) GetIdentityLookup(name string) *IdentityLookupDTO {
	if p.IdentityLookups == nil {
		return nil
	}

	for _, l := range *p.IdentityLookups {
		if l.GetName() == name {
			return &l
		}
	}

	return nil
}

func (p *InternalSaml2ServiceProviderDTO) AddIdentityLookup(name string) (IdentityLookupDTO, error) {

	// Initialize id lookup dto
	l := NewIdentityLookupDTO()
	l.SetName(name)
	l.AdditionalProperties = make(map[string]interface{})
	l.AdditionalProperties["@c"] = ".IdentityLookupDTO"

	var ls []IdentityLookupDTO

	if p.IdentityLookups == nil {
		ls = make([]IdentityLookupDTO, 0)
	} else {
		if p.GetIdentityLookup(name) != nil {
			return *l, fmt.Errorf("name already in use for identity lookup %s", name)
		}
		ls = *p.IdentityLookups
	}

	// Add a new element to : p.IdentityLookups
	ls = append(ls, *l)

	p.IdentityLookups = &ls

	return *l, nil
}

// Return true if the element was deleted, false otherwise
func (p *InternalSaml2ServiceProviderDTO) RemoveIdentityLookup(name string) bool {
	// Remove an element from : p.IdentityLookups

	if p.IdentityLookups == nil {
		return false
	}

	ls := *p.IdentityLookups
	var newLs []IdentityLookupDTO
	deleted := false

	for _, l := range ls {

		if l.GetName() == name {
			deleted = true
			continue
		}

		newLs = append(newLs, l)

	}

	p.IdentityLookups = &newLs

	return deleted
}
