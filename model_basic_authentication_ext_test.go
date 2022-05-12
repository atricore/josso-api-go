package jossoappi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAuthnMechanism(t *testing.T) {

	ba := NewBasicAuthenticationDTOInit()
	// Some defaults
	ba.SetName("basic-authn")
	ba.SetDisplayName("basic authn")
	ba.SetEnabled(false)

	ba.SetSimpleAuthnSaml2AuthnCtxClass("urn:oasis:names:tc:SAML:2.0:ac:classes:Password")
	ba.SetHashAlgorithm("SHA-256")
	ba.SetHashEncoding("BASE64")
	ba.SetSaltLength(2)

	ba.SetSaltPrefix("sp")
	ba.SetSaltSuffix("ss")

	m, err := ba.ToAuthnMechansim()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	baTest, err := m.ToBasicAuthn()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assert.Equal(t, ba.GetName(), baTest.GetName())
	assert.Equal(t, ba.GetDisplayName(), baTest.GetDisplayName())
	assert.Equal(t, ba.GetEnabled(), baTest.GetEnabled())
	assert.Equal(t, ba.GetSimpleAuthnSaml2AuthnCtxClass(), baTest.GetSimpleAuthnSaml2AuthnCtxClass())
	assert.Equal(t, ba.GetHashAlgorithm(), baTest.GetHashAlgorithm())
	assert.Equal(t, ba.GetHashEncoding(), baTest.GetHashEncoding())
	assert.Equal(t, ba.GetSaltLength(), baTest.GetSaltLength())
	assert.Equal(t, ba.GetSaltPrefix(), baTest.GetSaltPrefix())
	assert.Equal(t, ba.GetSaltSuffix(), baTest.GetSaltSuffix())

}
