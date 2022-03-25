package jossoappi

import (
	"fmt"
	"testing"
)

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
