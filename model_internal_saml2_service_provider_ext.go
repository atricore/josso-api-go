package jossoappi

func (i *InternalSaml2ServiceProviderDTO) GetSamlR2SPConfig() *SamlR2SPConfigDTO {

	c := i.GetConfig()
	var saml2sp SamlR2SPConfigDTO
	saml2sp.SetDescription(c.GetDescription())
	saml2sp.SetDisplayName(c.GetDisplayName())
	saml2sp.SetElementId(c.GetElementId())
	saml2sp.SetId(c.GetId())
	saml2sp.SetName(c.GetName())

	KeystoreDTO := toKeyStoreDTO((c.AdditionalProperties["keyStoreDTO"].(map[string]interface{})))
	saml2sp.SetSigner(*KeystoreDTO)
	saml2sp.SetUseSampleStore(c.AdditionalProperties["useSampleStore"].(bool))
	saml2sp.SetUseSystemStore(c.AdditionalProperties["useSystemStore"].(bool))

	return &saml2sp

}

func toKeyStoremap(dto KeystoreDTO) *map[string]interface{} {

	props := make(map[string]interface{})

	props["certificateAlias"] = dto.GetCertificateAlias()
	props["displayName"] = dto.GetDisplayName()
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["keystorePassOnly"] = dto.GetKeystorePassOnly()
	props["name"] = dto.GetName()
	props["password"] = dto.GetPassword()
	props["privateKeyName"] = dto.GetPrivateKeyName()
	props["privateKeyPassword"] = dto.GetPrivateKeyPassword()
	props["store"] = dto.GetStore()
	props["type"] = dto.GetType()

	return &props
}

// Transforms a map into an KeyStoreDTO
func toKeyStoreDTO(props map[string]interface{}) *KeystoreDTO {
	var storeProps map[string]interface{}
	Key := NewKeystoreDTO()
	Key.AdditionalProperties = make(map[string]interface{})

	//¿?Key.AdditionalProperties["@id"] = storeId
	Key.SetCertificateAlias((props["certificateAlias"].(string)))
	Key.SetDisplayName((props["displayName"].(string)))
	Key.SetElementId((props["elementId"].(string)))
	Key.SetId((props["id"].(int64)))
	Key.SetKeystorePassOnly((props["keystorePassOnly"].(bool)))
	Key.SetName((props["name"].(string)))
	Key.SetPassword((props["password"].(string)))
	Key.SetPrivateKeyName((props["privateKeyName"].(string)))
	Key.SetPrivateKeyPassword((props["privateKeyPassword"].(string)))
	Key.SetPrivateKeyPassword((props["privateKeyPassword"].(string)))

	Key.SetType((props["type"].(string)))

	resourceProps := storeProps["store"].(map[string]interface{})
	Key.Store = NewResourceDTO()
	Key.Store.AdditionalProperties = map[string]interface{}{}
	Key.Store.AdditionalProperties["@id"] = resourceProps["@id"]
	Key.Store.DisplayName = IPtrString(resourceProps["displayName"])
	Key.Store.ElementId = IPtrString(resourceProps["elementId"])
	Key.Store.Name = IPtrString(resourceProps["name"])
	Key.Store.Uri = IPtrString(resourceProps["uri"])

	return Key
}

func (i *InternalSaml2ServiceProviderDTO) SetProviderConfig(saml2sp *SamlR2SPConfigDTO) {

	var c SamlR2SPConfigDTO

	c.SetDescription(saml2sp.GetDescription())
	c.SetDisplayName(saml2sp.GetDisplayName())
	c.SetElementId(saml2sp.GetElementId())
	c.SetId(saml2sp.GetId())
	c.SetName(saml2sp.GetName())
	c.SetSigner(saml2sp.GetSigner())
	c.SetUseSampleStore(saml2sp.GetUseSampleStore())
	c.SetUseSystemStore(saml2sp.GetUseSystemStore())
	c.AdditionalProperties["keyStore"] = toKeyStoremap(c.GetSigner())

	
}
