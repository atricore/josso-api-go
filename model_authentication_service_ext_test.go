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

func TestToClientCertAuthnSvc(t *testing.T) {
	cas := NewClientCertAuthnSvcDTOInit()

	cas.SetId(10)
	cas.SetElementId("_942978F0-071D-428F-94EC-D2396D2B491A")
	cas.SetName("ldap-authn-1")
	cas.SetDisplayName("ldap authn 1")
	cas.SetDescription("ldap authn 1")

	cas.SetClrEnabled(false)
	cas.SetCrlRefreshSeconds(0)
	cas.SetCrlUrl("")
	cas.SetOcspEnabled(false)
	cas.SetOcspServer("")
	cas.SetOcspserver("")
	cas.SetUid("")

	m, err := cas.toClientCertAuthSvc()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	casTest, err := m.toClientCertAuthnSvc()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assert.Equal(t, cas.GetId(), casTest.GetId())
	assert.Equal(t, cas.GetElementId(), casTest.GetElementId())
	assert.Equal(t, cas.GetName(), casTest.GetName())
	assert.Equal(t, cas.GetDisplayName(), casTest.GetDisplayName())
	assert.Equal(t, cas.GetDescription(), casTest.GetDescription())

	assert.Equal(t, cas.GetClrEnabled(), casTest.GetClrEnabled())
	assert.Equal(t, cas.GetCrlRefreshSeconds(), casTest.GetCrlRefreshSeconds())
	assert.Equal(t, cas.GetCrlUrl(), casTest.GetCrlUrl())
	assert.Equal(t, cas.GetOcspEnabled(), casTest.GetOcspEnabled())
	assert.Equal(t, cas.GetOcspServer(), casTest.GetOcspServer())
	assert.Equal(t, cas.GetOcspserver(), casTest.GetOcspserver())
	assert.Equal(t, cas.GetUid(), casTest.GetUid())
}

func TestToWindowsIntegratedAuthn(t *testing.T) {
	wia := NewWindowsintegratedAuthnDTOInit()

	wia.SetId(10)
	wia.SetElementId("_942978F0-071D-428F-94EC-D2396D2B491A")
	wia.SetName("ldap-authn-1")
	wia.SetDisplayName("ldap authn 1")
	wia.SetDescription("ldap authn 1")

	wia.SetDomain("")
	wia.SetDomainController("")
	wia.SetHost("")
	wia.SetOverwriteKerberosSetup(false)
	wia.SetPort(0)
	wia.SetProtocol("")
	wia.SetServiceClass("")
	wia.SetServiceName("")

	m, err := wia.toWindowsIntegratedAuth()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	wiaTest, err := m.toWindowsIntegratedAuthn()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assert.Equal(t, wia.GetId(), wiaTest.GetId())
	assert.Equal(t, wia.GetElementId(), wiaTest.GetElementId())
	assert.Equal(t, wia.GetName(), wiaTest.GetName())
	assert.Equal(t, wia.GetDisplayName(), wiaTest.GetDisplayName())
	assert.Equal(t, wia.GetDescription(), wiaTest.GetDescription())

	assert.Equal(t, wia.GetDomain(), wiaTest.GetDomain())
	assert.Equal(t, wia.GetDomainController(), wiaTest.GetDomainController())
	assert.Equal(t, wia.GetHost(), wiaTest.GetHost())
	assert.Equal(t, wia.GetOverwriteKerberosSetup(), wiaTest.GetOverwriteKerberosSetup())
	assert.Equal(t, wia.GetPort(), wiaTest.GetPort())
	assert.Equal(t, wia.GetProtocol(), wiaTest.GetProtocol())
	assert.Equal(t, wia.GetServiceClass(), wiaTest.GetServiceClass())
	assert.Equal(t, wia.GetServiceName(), wiaTest.GetServiceName())
}

func TestToOauth2PreAuthnSvc(t *testing.T) {
	oaut2 := NewOAuth2PreAuthenticationServiceDTO()

	oaut2.SetId(0)
	oaut2.SetElementId("")
	oaut2.SetName("")
	oaut2.SetDisplayName("")
	oaut2.SetDescription("")

	oaut2.SetAuthnService("")
	oaut2.SetExternalAuth(false)
	oaut2.SetRememberMe(false)

	m, err := oaut2.toOauth2PreAuthnSvc()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	oaut2Test, err := m.toOauth2PreAuthnSvs()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assert.Equal(t, oaut2.GetId(), oaut2Test.GetId())
	assert.Equal(t, oaut2.GetElementId(), oaut2Test.GetElementId())
	assert.Equal(t, oaut2.GetName(), oaut2Test.GetName())
	assert.Equal(t, oaut2.GetDisplayName(), oaut2Test.GetDisplayName())
	assert.Equal(t, oaut2.GetDescription(), oaut2Test.GetDescription())

	assert.Equal(t, oaut2.GetAuthnService(), oaut2Test.GetAuthnService())
	assert.Equal(t, oaut2.GetExternalAuth(), oaut2Test.GetExternalAuth())
	assert.Equal(t, oaut2.GetRememberMe(), oaut2Test.GetRememberMe())
}
