package jossoappi

func toKeyStoreMap(store KeystoreDTO) map[string]interface{} {

	storeProps := make(map[string]interface{})

	storeProps["@id"] = store.AdditionalProperties["@id"].(int)
	storeProps["certificateAlias"] = store.GetCertificateAlias
	storeProps["displayName"] = store.GetDisplayName()
	storeProps["elementId"] = store.GetElementId()
	storeProps["id"] = store.GetId()
	storeProps["keystorePassOnly"] = store.GetKeystorePassOnly()
	storeProps["name"] = store.GetName()
	storeProps["password"] = store.GetPassword()
	storeProps["privateKeyName"] = store.GetPrivateKeyName()
	storeProps["privateKeyPassword"] = store.GetPrivateKeyPassword()
	storeProps["type"] = store.GetType()

	resourceProps := make(map[string]interface{})
	storeProps["store"] = resourceProps

	resourceProps["@id"] = store.Store.AdditionalProperties["@id"]
	resourceProps["displayName"] = store.Store.GetDisplayName()
	resourceProps["elementId"] = store.Store.GetElementId()
	resourceProps["name"] = store.Store.GetName()
	resourceProps["uri"] = store.Store.GetUri()
	resourceProps["value"] = store.Store.GetValue()

	return storeProps

}

// Transforms a map into an KeyStoreDTO
func toKeyStoreDTO(storeId int, props map[string]interface{}) *KeystoreDTO {

	var storeProps map[string]interface{}

	store := NewKeystoreDTO()
	store.AdditionalProperties = make(map[string]interface{})

	store.AdditionalProperties["@id"] = storeId
	store.SetCertificateAlias((props["certificateAlias"].(string)))
	store.SetDisplayName((props["displayName"].(string)))
	store.SetElementId((props["elementId"].(string)))
	store.SetId((props["id"].(int64)))
	store.SetKeystorePassOnly((props["keystorePassOnly"].(bool)))
	store.SetName((props["name"].(string)))
	store.SetPassword((props["password"].(string)))
	store.SetPrivateKeyName((props["privateKeyName"].(string)))
	store.SetPrivateKeyPassword((props["privateKeyPassword"].(string)))
	store.SetPrivateKeyPassword((props["privateKeyPassword"].(string)))
	store.SetType((props["type"].(string)))

	resourceProps := storeProps["store"].(map[string]interface{})
	store.Store = NewResourceDTO()
	store.Store.AdditionalProperties = map[string]interface{}{}
	store.Store.AdditionalProperties["@id"] = resourceProps["@id"]
	store.Store.DisplayName = IPtrString(resourceProps["displayName"])
	store.Store.ElementId = IPtrString(resourceProps["elementId"])
	store.Store.Name = IPtrString(resourceProps["name"])
	store.Store.Uri = IPtrString(resourceProps["uri"])
	store.Store.Value = IPtrString(resourceProps["value"])

	return store
}

func NewKeystoreDTOWithOK() *KeystoreDTO {
	this := KeystoreDTO{}
	this.AdditionalProperties = make(map[string]interface{})
	this.AdditionalProperties["@c"] = ".KeystoreDTO"

	this.Store = NewResourceDTO()
	this.Store.AdditionalProperties = map[string]interface{}{}
	this.Store.AdditionalProperties["@c"] = ".ResourceDTO"

	return &this
}
