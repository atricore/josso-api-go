# CreateIdPReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**SamlR2IDPConfigDTO**](SamlR2IDPConfigDTO.md) |  | [optional] 
**IdaName** | Pointer to **string** |  | [optional] 
**Idp** | Pointer to [**IdentityProviderDTO**](IdentityProviderDTO.md) |  | [optional] 

## Methods

### NewCreateIdPReq

`func NewCreateIdPReq() *CreateIdPReq`

NewCreateIdPReq instantiates a new CreateIdPReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdPReqWithDefaults

`func NewCreateIdPReqWithDefaults() *CreateIdPReq`

NewCreateIdPReqWithDefaults instantiates a new CreateIdPReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateIdPReq) GetConfig() SamlR2IDPConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateIdPReq) GetConfigOk() (*SamlR2IDPConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateIdPReq) SetConfig(v SamlR2IDPConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateIdPReq) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetIdaName

`func (o *CreateIdPReq) GetIdaName() string`

GetIdaName returns the IdaName field if non-nil, zero value otherwise.

### GetIdaNameOk

`func (o *CreateIdPReq) GetIdaNameOk() (*string, bool)`

GetIdaNameOk returns a tuple with the IdaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdaName

`func (o *CreateIdPReq) SetIdaName(v string)`

SetIdaName sets IdaName field to given value.

### HasIdaName

`func (o *CreateIdPReq) HasIdaName() bool`

HasIdaName returns a boolean if a field has been set.

### GetIdp

`func (o *CreateIdPReq) GetIdp() IdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *CreateIdPReq) GetIdpOk() (*IdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *CreateIdPReq) SetIdp(v IdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *CreateIdPReq) HasIdp() bool`

HasIdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


