package jossoappi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSPKeystore(t *testing.T) {

	// 1. create empty idp dto
	p := NewInternalSaml2ServiceProviderDTO()
	p.AdditionalProperties = make(map[string]interface{})

	// 2. Create IDP config
	spCfg := NewSamlR2SPConfigDTOInit()

	// Set KS value

	s := NewResourceDTOInit("signer keystore", "pkcs#12 signer keystore", "asdfasdf")

	ks := NewKeystoreDTOInit("signer", "pkcs#12 signer", s)
	ks.SetPassword("1234")
	ks.SetKeystorePassOnly(true)
	ks.SetStore(*s)

	spCfg.SetSigner(*ks)
	spCfg.SetEncrypter(*ks)

	err := p.SetSamlR2SPConfig(spCfg)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	idpCfgTest, err := p.GetSamlR2SPConfig()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	ksTest := idpCfgTest.GetSigner()
	sTest := ksTest.GetStore()

	assert.Equal(t, spCfg.GetDescription(), idpCfgTest.GetDescription(), "description does not match")
	assert.Equal(t, ks.GetPassword(), ksTest.GetPassword(), "password does not match")
	assert.Equal(t, ks.GetCertificateAlias(), ksTest.GetCertificateAlias(), "certificate alias does not match")
	assert.Equal(t, ks.GetKeystorePassOnly(), ksTest.GetKeystorePassOnly(), "keystore pass only does not match")
	assert.Equal(t, ks.GetPrivateKeyName(), ksTest.GetPrivateKeyName(), "private key name pass only does not match")
	assert.Equal(t, ks.GetPrivateKeyPassword(), ksTest.GetPrivateKeyPassword(), "private password name pass only does not match")
	assert.Equal(t, s.GetValue(), sTest.GetValue(), "keystore value does not match")

	// 3. Convert to provider config

}
