package jossoappi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDirectoryAuthnSvc(t *testing.T) {
	das := NewDirectoryAuthnSvcDTOInit()

	das.SetId(10)
	das.SetElementId("_942978F0-071D-428F-94EC-D2396D2B491A")
	das.SetName("ldap-authn-1")
	das.SetDisplayName("ldap authn 1")
	das.SetDescription("ldap authn 1")

	das.SetInitialContextFactory("com.sun.jndi.ldap.LdapCtxFactory")
	das.SetProviderUrl("ldap://localhost:389")
	das.SetPerformDnSearch(false)
	das.SetPasswordPolicy("none")
	das.SetSecurityAuthentication("simple")
	das.SetUsersCtxDN("ou=People,dc=my-domain,dc=com")
	das.SetPrincipalUidAttributeID("uid")
	das.SetSecurityPrincipal("uid=admin,ou=system")
	das.SetSecurityCredential("changeme")
	das.SetLdapSearchScope("subtree")
	das.SetSimpleAuthnSaml2AuthnCtxClass("urn:oasis:names:tc:SAML:2.0:ac:classes:Password")
	das.SetReferrals("follow")
	das.SetIncludeOperationalAttributes(false)

	m, err := das.toAuthnSvc()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	dasTest, err := m.toDirectoryAuthnSvc()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assert.Equal(t, das.GetId(), dasTest.GetId())
	assert.Equal(t, das.GetElementId(), dasTest.GetElementId())
	assert.Equal(t, das.GetName(), dasTest.GetName())
	assert.Equal(t, das.GetDisplayName(), dasTest.GetDisplayName())
	assert.Equal(t, das.GetDescription(), dasTest.GetDescription())
	
	assert.Equal(t, das.GetInitialContextFactory(), dasTest.GetInitialContextFactory())
	assert.Equal(t, das.GetProviderUrl(), dasTest.GetProviderUrl())
	assert.Equal(t, das.GetPerformDnSearch(), dasTest.GetPerformDnSearch())
	assert.Equal(t, das.GetPasswordPolicy(), dasTest.GetPasswordPolicy())
	assert.Equal(t, das.GetSecurityAuthentication(), dasTest.GetSecurityAuthentication())
	assert.Equal(t, das.GetUsersCtxDN(), dasTest.GetUsersCtxDN())
	assert.Equal(t, das.GetPrincipalUidAttributeID(), dasTest.GetPrincipalUidAttributeID())
	assert.Equal(t, das.GetSecurityPrincipal(), dasTest.GetSecurityPrincipal())
	assert.Equal(t, das.GetSecurityCredential(), dasTest.GetSecurityCredential())
	assert.Equal(t, das.GetLdapSearchScope(), dasTest.GetLdapSearchScope())
	assert.Equal(t, das.GetSimpleAuthnSaml2AuthnCtxClass(), dasTest.GetSimpleAuthnSaml2AuthnCtxClass())
	assert.Equal(t, das.GetReferrals(), dasTest.GetReferrals())
	assert.Equal(t, das.GetIncludeOperationalAttributes(), dasTest.GetIncludeOperationalAttributes())
}
