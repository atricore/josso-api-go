# GetTypesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idpc** | Pointer to [**IdentityProviderChannelDTO**](IdentityProviderChannelDTO.md) |  | [optional] 
**Spc** | Pointer to [**InternalSaml2ServiceProviderChannelDTO**](InternalSaml2ServiceProviderChannelDTO.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

