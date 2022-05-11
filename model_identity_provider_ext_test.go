package jossoappi

import (
	"fmt"
	"testing"
)

func TestKeystoreAdd(t *testing.T) {

	// 1. create empty idp dto
	p := NewIdentityProviderDTO()
	p.AdditionalProperties = make(map[string]interface{})

	// 2. Create IDP config
	idpCfg := NewSamlR2IDPConfigDTOWithOK()

	// Set KS value
	ks := NewKeystoreDTOWithOK()
	s := ks.GetStore()
	s.SetValue("asdf")
	ks.SetPassword("1234")
	ks.SetKeystorePassOnly(true)

	idpCfg.SetSigner(*ks)
	idpCfg.SetEncrypter(*ks)

	err := p.SetSamlR2IDPConfig(idpCfg)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	idpCfgTest, err := p.GetSamlR2IDPConfig()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	
	if idpCfgTest.GetDescription() != idpCfg.GetDescription()
	// 3. Convert to provider config

}

func TestAddIdLookup(t *testing.T) {

	// 1. create empty idp dto
	p := NewIdentityProviderDTO()
	p.AdditionalProperties = make(map[string]interface{})

	// 2. add id lookup
	_, err := p.AddIdentityLookup("ilk-1")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	l, err := p.AddIdentityLookup("ilk-2")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	for _, l = range p.IdentityLookups {
		fmt.Printf("%s\n", l.GetName())
	}

	// 3. verify id lookups size and names
	if len(p.IdentityLookups) != 2 {
		t.Errorf("Invalid number of identity lookups %d", len(p.IdentityLookups))
		return
	}

}

func TestRemoveIdLookup(t *testing.T) {

	// 1. create empty idp dto
	p := NewIdentityProviderDTO()
	p.AdditionalProperties = make(map[string]interface{})

	// 2. add id lookup
	p.AddIdentityLookup("ilk-1")
	p.AddIdentityLookup("ilk-2")

	// 3. verify id lookups size and names
	if len(p.IdentityLookups) != 2 {
		t.Errorf("Invalid number of identity lookups %d", len(p.IdentityLookups))
		return
	}

	// 4. remove id lookup

	d := p.RemoveIdentityLookup("ilk-1")

	if !d {
		t.Errorf("Sould have deleted ilk-1")
		return
	}

	if len(p.IdentityLookups) != 1 {
		t.Errorf("Invalid number of identity lookups %d", len(p.IdentityLookups))
		return
	}

	if (p.IdentityLookups)[0].GetName() != "ilk-2" {
		t.Errorf("Invalid name %s", (p.IdentityLookups)[0].GetName())
		return

	}

	d = p.RemoveIdentityLookup("ilk-2")
	if !d {
		t.Errorf("Sould have deleted ilk-2")
		return
	}
	if len(p.IdentityLookups) != 0 {
		t.Errorf("Invalid number of identity lookups %d", len(p.IdentityLookups))
		return
	}

	d = p.RemoveIdentityLookup("asfadsf")
	if d {
		t.Errorf("Sould have not deleted !")
		return
	}

	// 5. verify id lookups size and names

}
