package jossoappi

import "fmt"

func (cfg *ProviderConfigDTO) GetStoreProps() (map[string]interface{}, int, error) {
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
