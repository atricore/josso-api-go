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

func (m AuthenticationMechanismDTO) IsDirectoryAuthnSvc() bool {
	return m.AdditionalProperties["@c"] == ".DirectoryAuthenticationServiceDTO"
}
