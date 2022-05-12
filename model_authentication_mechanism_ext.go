package jossoappi

import "fmt"

func (m AuthenticationMechanismDTO) ToBasicAuthn() (*BasicAuthenticationDTO, error) {
	ba := NewBasicAuthenticationWithOK()

	if m.AdditionalProperties["@c"] != ba.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", m.AdditionalProperties["@c"])
	}

	ba.SetName(m.GetName())
	ba.SetDisplayName(m.GetDisplayName())
	ba.SetPriority(m.GetPriority())
	ba.SetEnabled(AsBool(m.AdditionalProperties["enabled"], false))

	ba.SetHashAlgorithm(m.AdditionalProperties["hashAlgorithm"].(string))
	ba.SetHashEncoding(m.AdditionalProperties["hashEncoding"].(string))
	ba.SetIgnoreUsernameCase(AsBool(m.AdditionalProperties["ignoreUsernamecase"], false))

	//ba.SetIgnorePasswordCase(m.AdditionalProperties["ignorePassowordCase"].(bool))

	ba.SetSaltLength(AsInt32(m.AdditionalProperties["saltLength"], 0))
	ba.SetSaltPrefix(AsString(m.AdditionalProperties["saltPrefix"], ""))
	ba.SetSaltSuffix(AsString(m.AdditionalProperties["saltSuffix"], ""))
	//authn.AdditionalProperties["impersonateUserPolicy"]
	ba.SetSimpleAuthnSaml2AuthnCtxClass(m.AdditionalProperties["simpleAuthnSaml2AuthnCtxClass"].(string))

	return ba, nil
}

func (m AuthenticationMechanismDTO) IsBasicAuthn() bool {
	return m.AdditionalProperties["@c"] == ".BasicAuthenticationDTO"
}
