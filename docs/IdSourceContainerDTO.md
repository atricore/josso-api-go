# IdSourceContainerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdSource** | Pointer to [**IdentitySourceDTO**](IdentitySourceDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewIdSourceContainerDTO

`func NewIdSourceContainerDTO() *IdSourceContainerDTO`

NewIdSourceContainerDTO instantiates a new IdSourceContainerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdSourceContainerDTOWithDefaults

`func NewIdSourceContainerDTOWithDefaults() *IdSourceContainerDTO`

NewIdSourceContainerDTOWithDefaults instantiates a new IdSourceContainerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdSource

`func (o *IdSourceContainerDTO) GetIdSource() IdentitySourceDTO`

GetIdSource returns the IdSource field if non-nil, zero value otherwise.

### GetIdSourceOk

`func (o *IdSourceContainerDTO) GetIdSourceOk() (*IdentitySourceDTO, bool)`

GetIdSourceOk returns a tuple with the IdSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSource

`func (o *IdSourceContainerDTO) SetIdSource(v IdentitySourceDTO)`

SetIdSource sets IdSource field to given value.

### HasIdSource

`func (o *IdSourceContainerDTO) HasIdSource() bool`

HasIdSource returns a boolean if a field has been set.

### GetName

`func (o *IdSourceContainerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdSourceContainerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdSourceContainerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdSourceContainerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IdSourceContainerDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdSourceContainerDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdSourceContainerDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdSourceContainerDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


