# GroupDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ExtraAttributes** | Pointer to [**[]AttributeValueDTO**](AttributeValueDTO.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupDTO

`func NewGroupDTO() *GroupDTO`

NewGroupDTO instantiates a new GroupDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDTOWithDefaults

`func NewGroupDTOWithDefaults() *GroupDTO`

NewGroupDTOWithDefaults instantiates a new GroupDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtraAttributes

`func (o *GroupDTO) GetExtraAttributes() []AttributeValueDTO`

GetExtraAttributes returns the ExtraAttributes field if non-nil, zero value otherwise.

### GetExtraAttributesOk

`func (o *GroupDTO) GetExtraAttributesOk() (*[]AttributeValueDTO, bool)`

GetExtraAttributesOk returns a tuple with the ExtraAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraAttributes

`func (o *GroupDTO) SetExtraAttributes(v []AttributeValueDTO)`

SetExtraAttributes sets ExtraAttributes field to given value.

### HasExtraAttributes

`func (o *GroupDTO) HasExtraAttributes() bool`

HasExtraAttributes returns a boolean if a field has been set.

### GetId

`func (o *GroupDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupDTO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


