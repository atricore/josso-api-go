package jossoappi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuthn(t *testing.T) {

	// 1. create empty idp dto
	p := NewIdentityProviderDTO()
	p.AdditionalProperties = make(map[string]interface{})

	ba := newBasicAuth(1)

	var bas []*BasicAuthenticationDTO
	bas = append(bas, ba)
	p.AddBasicAuthns(bas)

	assert.Equal(t, 1, len(p.AuthenticationMechanisms))

	basTest, err := p.GetBasicAuthns()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assert.Equal(t, len(bas), len(basTest), "invalid number of basic authentication mechanisms")

	baTest := basTest[0]

	assertEqual(t, ba, baTest)

}

func TestBasicAuthns(t *testing.T) {

	ba1 := newBasicAuth(1)
	ba2 := newBasicAuth(2)

	// 1. create empty idp dto
	p := NewIdentityProviderDTO()
	p.AdditionalProperties = make(map[string]interface{})

	var bas []*BasicAuthenticationDTO
	bas = append(bas, ba1)
	bas = append(bas, ba2)
	p.AddBasicAuthns(bas)
	assert.Equal(t, 2, len(p.AuthenticationMechanisms))

	basTest, err := p.GetBasicAuthns()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	for _, ba := range bas {
		f := false
		for _, baTest := range basTest {
			if ba.GetName() == baTest.GetName() {
				assertEqual(t, ba, baTest)
				f = true
			}
		}
		if !f {
			t.Errorf("%s authn mechanism not found in idp", ba.GetName())
		}

	}

}

func TestKeystore(t *testing.T) {

	// 1. create empty idp dto
	p := NewIdentityProviderDTO()
	p.AdditionalProperties = make(map[string]interface{})

	// 2. Create IDP config
	idpCfg := NewSamlR2IDPConfigDTOInit()

	// Set KS value

	s := NewResourceDTOInit("signer keystore", "pkcs#12 signer keystore", "asdfasdf")

	ks := NewKeystoreDTOInit("signer", "pkcs#12 signer", s)
	ks.SetPassword("1234")
	ks.SetKeystorePassOnly(true)
	ks.SetStore(*s)

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

	ksTest := idpCfgTest.GetSigner()
	sTest := ksTest.GetStore()

	assert.Equal(t, idpCfg.GetDescription(), idpCfgTest.GetDescription(), "description does not match")
	assert.Equal(t, ks.GetPassword(), ksTest.GetPassword(), "password does not match")
	assert.Equal(t, ks.GetCertificateAlias(), ksTest.GetCertificateAlias(), "certificate alias does not match")
	assert.Equal(t, ks.GetKeystorePassOnly(), ksTest.GetKeystorePassOnly(), "keystore pass only does not match")
	assert.Equal(t, ks.GetPrivateKeyName(), ksTest.GetPrivateKeyName(), "private key name pass only does not match")
	assert.Equal(t, ks.GetPrivateKeyPassword(), ksTest.GetPrivateKeyPassword(), "private password name pass only does not match")
	assert.Equal(t, s.GetValue(), sTest.GetValue(), "keystore value does not match")

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

func newBasicAuth(idx int32) *BasicAuthenticationDTO {
	ba := NewBasicAuthenticationDTOInit()

	ba.SetPriority(3 + idx)
	ba.SetSaltLength(2 + idx)
	ba.SetEnabled(true)

	ba.SetName(fmt.Sprintf("basic-authn-%d", idx))
	ba.SetDisplayName(fmt.Sprintf("Basic Authn-%d", idx))
	ba.SetHashAlgorithm(fmt.Sprintf("SHA-256-%d", idx))
	ba.SetHashEncoding(fmt.Sprintf("BASE64-%d", idx))
	ba.SetSaltSuffix(fmt.Sprintf("su-%d", idx))
	ba.SetSaltPrefix(fmt.Sprintf("sp-%d", idx))
	ba.SetSimpleAuthnSaml2AuthnCtxClass(fmt.Sprintf("SAML-CLASS-%d", idx))
	return ba
}

func assertEqual(t *testing.T, ba *BasicAuthenticationDTO, baTest *BasicAuthenticationDTO) {
	assert.Equal(t, ba.AdditionalProperties, baTest.AdditionalProperties)
	assert.Equal(t, ba.GetName(), baTest.GetName())
	assert.Equal(t, ba.GetDisplayName(), baTest.GetDisplayName())
	assert.Equal(t, ba.GetEnabled(), baTest.GetEnabled())
	assert.Equal(t, ba.GetHashAlgorithm(), baTest.GetHashAlgorithm())
	assert.Equal(t, ba.GetSaltLength(), baTest.GetSaltLength())
	assert.Equal(t, ba.GetSaltSuffix(), baTest.GetSaltSuffix())
	assert.Equal(t, ba.GetSaltPrefix(), baTest.GetSaltPrefix())
	assert.Equal(t, ba.GetSimpleAuthnSaml2AuthnCtxClass(), baTest.GetSimpleAuthnSaml2AuthnCtxClass())

}
