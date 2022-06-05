# ProviderContainerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtOidcRp** | Pointer to [**ExternalOpenIDConnectRelayingPartyDTO**](ExternalOpenIDConnectRelayingPartyDTO.md) |  | [optional] 
**ExtSaml2sp** | Pointer to [**ExternalSaml2ServiceProviderDTO**](ExternalSaml2ServiceProviderDTO.md) |  | [optional] 
**Idp** | Pointer to [**IdentityProviderDTO**](IdentityProviderDTO.md) |  | [optional] 
**Saml2sp** | Pointer to [**InternalSaml2ServiceProviderDTO**](InternalSaml2ServiceProviderDTO.md) |  | [optional] 

## Methods

### NewProviderContainerDTO

`func NewProviderContainerDTO() *ProviderContainerDTO`

NewProviderContainerDTO instantiates a new ProviderContainerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderContainerDTOWithDefaults

`func NewProviderContainerDTOWithDefaults() *ProviderContainerDTO`

NewProviderContainerDTOWithDefaults instantiates a new ProviderContainerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtOidcRp

`func (o *ProviderContainerDTO) GetExtOidcRp() ExternalOpenIDConnectRelayingPartyDTO`

GetExtOidcRp returns the ExtOidcRp field if non-nil, zero value otherwise.

### GetExtOidcRpOk

`func (o *ProviderContainerDTO) GetExtOidcRpOk() (*ExternalOpenIDConnectRelayingPartyDTO, bool)`

GetExtOidcRpOk returns a tuple with the ExtOidcRp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtOidcRp

`func (o *ProviderContainerDTO) SetExtOidcRp(v ExternalOpenIDConnectRelayingPartyDTO)`

SetExtOidcRp sets ExtOidcRp field to given value.

### HasExtOidcRp

`func (o *ProviderContainerDTO) HasExtOidcRp() bool`

HasExtOidcRp returns a boolean if a field has been set.

### GetExtSaml2sp

`func (o *ProviderContainerDTO) GetExtSaml2sp() ExternalSaml2ServiceProviderDTO`

GetExtSaml2sp returns the ExtSaml2sp field if non-nil, zero value otherwise.

### GetExtSaml2spOk

`func (o *ProviderContainerDTO) GetExtSaml2spOk() (*ExternalSaml2ServiceProviderDTO, bool)`

GetExtSaml2spOk returns a tuple with the ExtSaml2sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtSaml2sp

`func (o *ProviderContainerDTO) SetExtSaml2sp(v ExternalSaml2ServiceProviderDTO)`

SetExtSaml2sp sets ExtSaml2sp field to given value.

### HasExtSaml2sp

`func (o *ProviderContainerDTO) HasExtSaml2sp() bool`

HasExtSaml2sp returns a boolean if a field has been set.

### GetIdp

`func (o *ProviderContainerDTO) GetIdp() IdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *ProviderContainerDTO) GetIdpOk() (*IdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *ProviderContainerDTO) SetIdp(v IdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *ProviderContainerDTO) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetSaml2sp

`func (o *ProviderContainerDTO) GetSaml2sp() InternalSaml2ServiceProviderDTO`

GetSaml2sp returns the Saml2sp field if non-nil, zero value otherwise.

### GetSaml2spOk

`func (o *ProviderContainerDTO) GetSaml2spOk() (*InternalSaml2ServiceProviderDTO, bool)`

GetSaml2spOk returns a tuple with the Saml2sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml2sp

`func (o *ProviderContainerDTO) SetSaml2sp(v InternalSaml2ServiceProviderDTO)`

SetSaml2sp sets Saml2sp field to given value.

### HasSaml2sp

`func (o *ProviderContainerDTO) HasSaml2sp() bool`

HasSaml2sp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


