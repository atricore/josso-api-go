# GetTypesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeMapperProfileDTO** | Pointer to [**AttributeMapperProfileDTO**](AttributeMapperProfileDTO.md) |  | [optional] 
**BasicAuthnMechanism** | Pointer to [**BasicAuthenticationDTO**](BasicAuthenticationDTO.md) |  | [optional] 
**BuiltInAttributeProfile** | Pointer to [**BuiltInAttributeProfileDTO**](BuiltInAttributeProfileDTO.md) |  | [optional] 
**Idpc** | Pointer to [**IdentityProviderChannelDTO**](IdentityProviderChannelDTO.md) |  | [optional] 
**Spc** | Pointer to [**InternalSaml2ServiceProviderChannelDTO**](InternalSaml2ServiceProviderChannelDTO.md) |  | [optional] 
**TotpAuthnSvc** | Pointer to [**TOTPAuthenticationServiceDTO**](TOTPAuthenticationServiceDTO.md) |  | [optional] 

## Methods

### NewGetTypesRes

`func NewGetTypesRes() *GetTypesRes`

NewGetTypesRes instantiates a new GetTypesRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTypesResWithDefaults

`func NewGetTypesResWithDefaults() *GetTypesRes`

NewGetTypesResWithDefaults instantiates a new GetTypesRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeMapperProfileDTO

`func (o *GetTypesRes) GetAttributeMapperProfileDTO() AttributeMapperProfileDTO`

GetAttributeMapperProfileDTO returns the AttributeMapperProfileDTO field if non-nil, zero value otherwise.

### GetAttributeMapperProfileDTOOk

`func (o *GetTypesRes) GetAttributeMapperProfileDTOOk() (*AttributeMapperProfileDTO, bool)`

GetAttributeMapperProfileDTOOk returns a tuple with the AttributeMapperProfileDTO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapperProfileDTO

`func (o *GetTypesRes) SetAttributeMapperProfileDTO(v AttributeMapperProfileDTO)`

SetAttributeMapperProfileDTO sets AttributeMapperProfileDTO field to given value.

### HasAttributeMapperProfileDTO

`func (o *GetTypesRes) HasAttributeMapperProfileDTO() bool`

HasAttributeMapperProfileDTO returns a boolean if a field has been set.

### GetBasicAuthnMechanism

`func (o *GetTypesRes) GetBasicAuthnMechanism() BasicAuthenticationDTO`

GetBasicAuthnMechanism returns the BasicAuthnMechanism field if non-nil, zero value otherwise.

### GetBasicAuthnMechanismOk

`func (o *GetTypesRes) GetBasicAuthnMechanismOk() (*BasicAuthenticationDTO, bool)`

GetBasicAuthnMechanismOk returns a tuple with the BasicAuthnMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthnMechanism

`func (o *GetTypesRes) SetBasicAuthnMechanism(v BasicAuthenticationDTO)`

SetBasicAuthnMechanism sets BasicAuthnMechanism field to given value.

### HasBasicAuthnMechanism

`func (o *GetTypesRes) HasBasicAuthnMechanism() bool`

HasBasicAuthnMechanism returns a boolean if a field has been set.

### GetBuiltInAttributeProfile

`func (o *GetTypesRes) GetBuiltInAttributeProfile() BuiltInAttributeProfileDTO`

GetBuiltInAttributeProfile returns the BuiltInAttributeProfile field if non-nil, zero value otherwise.

### GetBuiltInAttributeProfileOk

`func (o *GetTypesRes) GetBuiltInAttributeProfileOk() (*BuiltInAttributeProfileDTO, bool)`

GetBuiltInAttributeProfileOk returns a tuple with the BuiltInAttributeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInAttributeProfile

`func (o *GetTypesRes) SetBuiltInAttributeProfile(v BuiltInAttributeProfileDTO)`

SetBuiltInAttributeProfile sets BuiltInAttributeProfile field to given value.

### HasBuiltInAttributeProfile

`func (o *GetTypesRes) HasBuiltInAttributeProfile() bool`

HasBuiltInAttributeProfile returns a boolean if a field has been set.

### GetIdpc

`func (o *GetTypesRes) GetIdpc() IdentityProviderChannelDTO`

GetIdpc returns the Idpc field if non-nil, zero value otherwise.

### GetIdpcOk

`func (o *GetTypesRes) GetIdpcOk() (*IdentityProviderChannelDTO, bool)`

GetIdpcOk returns a tuple with the Idpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpc

`func (o *GetTypesRes) SetIdpc(v IdentityProviderChannelDTO)`

SetIdpc sets Idpc field to given value.

### HasIdpc

`func (o *GetTypesRes) HasIdpc() bool`

HasIdpc returns a boolean if a field has been set.

### GetSpc

`func (o *GetTypesRes) GetSpc() InternalSaml2ServiceProviderChannelDTO`

GetSpc returns the Spc field if non-nil, zero value otherwise.

### GetSpcOk

`func (o *GetTypesRes) GetSpcOk() (*InternalSaml2ServiceProviderChannelDTO, bool)`

GetSpcOk returns a tuple with the Spc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpc

`func (o *GetTypesRes) SetSpc(v InternalSaml2ServiceProviderChannelDTO)`

SetSpc sets Spc field to given value.

### HasSpc

`func (o *GetTypesRes) HasSpc() bool`

HasSpc returns a boolean if a field has been set.

### GetTotpAuthnSvc

`func (o *GetTypesRes) GetTotpAuthnSvc() TOTPAuthenticationServiceDTO`

GetTotpAuthnSvc returns the TotpAuthnSvc field if non-nil, zero value otherwise.

### GetTotpAuthnSvcOk

`func (o *GetTypesRes) GetTotpAuthnSvcOk() (*TOTPAuthenticationServiceDTO, bool)`

GetTotpAuthnSvcOk returns a tuple with the TotpAuthnSvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpAuthnSvc

`func (o *GetTypesRes) SetTotpAuthnSvc(v TOTPAuthenticationServiceDTO)`

SetTotpAuthnSvc sets TotpAuthnSvc field to given value.

### HasTotpAuthnSvc

`func (o *GetTypesRes) HasTotpAuthnSvc() bool`

HasTotpAuthnSvc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


