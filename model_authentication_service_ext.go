package jossoappi

import "fmt"

// AuthenticationServiceDTO -> DirectoryAuthenticationServiceDTO
func (svc AuthenticationServiceDTO) toDirectoryAuthnSvc() (*DirectoryAuthenticationServiceDTO, error) {
	das := NewDirectoryAuthnSvcDTOInit()

	if svc.AdditionalProperties["@c"] != das.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", das.AdditionalProperties["@c"])
	}

	das.SetId(svc.GetId())
	das.SetElementId(svc.GetElementId())
	das.SetName(svc.GetName())
	das.SetDisplayName(svc.GetDisplayName())
	das.SetDescription(svc.GetDescription())
	das.SetDelegatedAuthentications(svc.GetDelegatedAuthentications())

	das.SetInitialContextFactory(AsString(svc.AdditionalProperties["initialContextFactory"], ""))
	das.SetProviderUrl(AsString(svc.AdditionalProperties["providerUrl"], ""))
	das.SetPerformDnSearch(AsBool(svc.AdditionalProperties["performDnSearch"], false))
	das.SetPasswordPolicy(AsString(svc.AdditionalProperties["passwordPolicy"], ""))
	das.SetSecurityAuthentication(AsString(svc.AdditionalProperties["securityAuthentication"], ""))
	das.SetUsersCtxDN(AsString(svc.AdditionalProperties["usersCtxDN"], ""))
	das.SetPrincipalUidAttributeID(AsString(svc.AdditionalProperties["principalUidAttributeID"], ""))
	das.SetSecurityPrincipal(AsString(svc.AdditionalProperties["securityPrincipal"], ""))
	das.SetSecurityCredential(AsString(svc.AdditionalProperties["securityCredential"], ""))
	das.SetLdapSearchScope(AsString(svc.AdditionalProperties["ldapSearchScope"], ""))
	das.SetSimpleAuthnSaml2AuthnCtxClass(AsString(svc.AdditionalProperties["simpleAuthnSaml2AuthnCtxClass"], ""))
	das.SetReferrals(AsString(svc.AdditionalProperties["referrals"], ""))
	das.SetIncludeOperationalAttributes(AsBool(svc.AdditionalProperties["includeOperationalAttributes"], false))

	return das, nil
}

// AuthenticationServiceDTO -> ClientCertAuthnServiceDTO
func (svc AuthenticationServiceDTO) toClientCertAuthnSvc() (*ClientCertAuthnServiceDTO, error) {
	cas := NewClientCertAuthnSvcDTOInit()

	if svc.AdditionalProperties["@c"] != cas.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", cas.AdditionalProperties["@c"])
	}

	cas.SetId(svc.GetId())
	cas.SetElementId(svc.GetElementId())
	cas.SetName(svc.GetName())
	cas.SetDisplayName(svc.GetDisplayName())
	cas.SetDescription(svc.GetDescription())
	cas.SetDelegatedAuthentications(svc.GetDelegatedAuthentications())

	cas.SetClrEnabled(AsBool(svc.AdditionalProperties["clrEnabled"], false))
	cas.SetCrlRefreshSeconds(AsInt32(svc.AdditionalProperties["crlRefreshSeconds"], 0))
	cas.SetCrlUrl(AsString(svc.AdditionalProperties["crlUrl"], ""))
	cas.SetOcspEnabled(AsBool(svc.AdditionalProperties["ocspEnabled"], false))
	cas.SetOcspServer(AsString(svc.AdditionalProperties["ocspServer"], ""))
	cas.SetOcspserver(AsString(svc.AdditionalProperties["ocspserver"], ""))
	cas.SetUid(AsString(svc.AdditionalProperties["uid"], ""))

	return cas, nil
}

// AuthenticationServiceDTO -> WindowsIntegratedAuthenticationDTO
func (svc AuthenticationServiceDTO) toWindowsIntegratedAuthn() (*WindowsIntegratedAuthenticationDTO, error) {
	wia := NewWindowsintegratedAuthnDTOInit()

	if svc.AdditionalProperties["@c"] != wia.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", wia.AdditionalProperties["@c"])
	}
	wia.SetId(svc.GetId())
	wia.SetElementId(svc.GetElementId())
	wia.SetName(svc.GetName())
	wia.SetDisplayName(svc.GetDisplayName())
	wia.SetDescription(svc.GetDescription())
	wia.SetDelegatedAuthentications(svc.GetDelegatedAuthentications())

	// TODO : wia.SetKeyTab ?
	// ktMap := svc.AdditionalProperties["keytab"].(map[string]interface{})
	// ktValue := ktMap["value"]

	wia.SetDomain(AsString(svc.AdditionalProperties["domain"], ""))
	wia.SetDomainController(AsString(svc.AdditionalProperties["domainController"], ""))
	wia.SetHost(AsString(svc.AdditionalProperties["host"], ""))
	wia.SetOverwriteKerberosSetup(AsBool(svc.AdditionalProperties["overwriteKerberosSetup"], false))
	wia.SetPort(AsInt32(svc.AdditionalProperties["port"], 0))
	wia.SetProtocol(AsString(svc.AdditionalProperties["protocol"], ""))
	wia.SetServiceClass(AsString(svc.AdditionalProperties["serviceClass"], ""))
	wia.SetServiceName(AsString(svc.AdditionalProperties["serviceName"], ""))

	return wia, nil
}

// AuthenticationServiceDTO -> OAuth2PreAuthenticationServiceDTO
func (svc AuthenticationServiceDTO) toOauth2PreAuthnSvs() (*OAuth2PreAuthenticationServiceDTO, error) {
	oaut2 := NewOauth2PreAuthnSvcDTOInit()

	if svc.AdditionalProperties["@c"] != oaut2.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", oaut2.AdditionalProperties["@c"])
	}

	oaut2.SetId(svc.GetId())
	oaut2.SetElementId(svc.GetElementId())
	oaut2.SetName(svc.GetName())
	oaut2.SetDisplayName(svc.GetDisplayName())
	oaut2.SetDescription(svc.GetDescription())
	oaut2.SetDelegatedAuthentications(svc.GetDelegatedAuthentications())

	oaut2.SetAuthnService(AsString(svc.AdditionalProperties["authnService"], ""))
	oaut2.SetExternalAuth((AsBool(svc.AdditionalProperties["externalAuth"], false)))
	oaut2.SetRememberMe(AsBool(svc.AdditionalProperties["rememberMe"], false))

	return oaut2, nil
}

func (m AuthenticationServiceDTO) IsDirectoryAuthnSvs() bool {
	return m.AdditionalProperties["@c"] == ".DirectoryAuthenticationServiceDTO"
}

func (m AuthenticationServiceDTO) IsClientCertAuthnSvs() bool {
	return m.AdditionalProperties["@c"] == ".ClientCertAuthnServiceDTO"
}

func (m AuthenticationServiceDTO) IsWindowsIntegratedAuthn() bool {
	return m.AdditionalProperties["@c"] == ".WindowsIntegratedAuthenticationDTO"
}

func (m AuthenticationServiceDTO) IsOauth2PreAuthnSvc() bool {
	return m.AdditionalProperties["@c"] == ".OAuth2PreAuthenticationServiceDTO"
}

func (p *IdentityProviderDTO) GetDirectoryAuthnSvc() ([]*DirectoryAuthenticationServiceDTO, error) {

	das := make([]*DirectoryAuthenticationServiceDTO, 0)

	for _, m := range p.GetAuthenticationMechanisms() {
		da := m.GetDelegatedAuthentication()
		as := da.GetAuthnService()
		if as.IsDirectoryAuthnSvs() {
			a, err := as.toDirectoryAuthnSvc()
			if err != nil {
				return das, err
			}
			das = append(das, a)
		}
	}

	return das, nil
}

func (p *IdentityProviderDTO) GetClientCertAuthnSvs() ([]*ClientCertAuthnServiceDTO, error) {

	cas := make([]*ClientCertAuthnServiceDTO, 0)

	for _, m := range p.AuthenticationMechanisms() {
		if m.IsClientCertAuthnSvc() {
			ca, err := m.toClientCertAuthnSvc()
			if err != nil {
				return cas, err
			}
			cas = append(cas, ca)
		}
	}

	return cas, nil
}

func (p *IdentityProviderDTO) GetWindowsIntegratedAuthn() ([]*WindowsIntegratedAuthenticationDTO, error) {

	wia := make([]*WindowsIntegratedAuthenticationDTO, 0)

	for _, m := range p.AuthenticationMechanisms() {
		if m.IsWindowsIntegratedAuthn() {
			wa, err := m.toWindowsIntegratedAuthn()
			if err != nil {
				return wia, err
			}
			wia = append(wia, wa)
		}
	}

	return wia, nil
}

func (p *IdentityProviderDTO) GetOauth2PreAuthnSvs() ([]*OAuth2PreAuthenticationServiceDTO, error) {

	oaut2 := make([]*OAuth2PreAuthenticationServiceDTO, 0)

	for _, m := range p.AuthenticationMechanisms() {
		if m.IsOauth2PreAuthnSvc() {
			oaut2svc, err := m.toOauth2PreAuthnSvc()
			if err != nil {
				return oaut2, err
			}
			oaut2 = append(oaut2, oaut2svc)
		}
	}

	return oaut2, nil
}
