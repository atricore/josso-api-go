# AttributeMapperProfileDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeMaps** | Pointer to [**[]AttributeMappingDTO**](AttributeMappingDTO.md) |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IncludeNonMappedProperties** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProfileType** | Pointer to **string** |  | [optional] 

## Methods

### NewAttributeMapperProfileDTO

`func NewAttributeMapperProfileDTO() *AttributeMapperProfileDTO`

NewAttributeMapperProfileDTO instantiates a new AttributeMapperProfileDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeMapperProfileDTOWithDefaults

`func NewAttributeMapperProfileDTOWithDefaults() *AttributeMapperProfileDTO`

NewAttributeMapperProfileDTOWithDefaults instantiates a new AttributeMapperProfileDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeMaps

`func (o *AttributeMapperProfileDTO) GetAttributeMaps() []AttributeMappingDTO`

GetAttributeMaps returns the AttributeMaps field if non-nil, zero value otherwise.

### GetAttributeMapsOk

`func (o *AttributeMapperProfileDTO) GetAttributeMapsOk() (*[]AttributeMappingDTO, bool)`

GetAttributeMapsOk returns a tuple with the AttributeMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMaps

`func (o *AttributeMapperProfileDTO) SetAttributeMaps(v []AttributeMappingDTO)`

SetAttributeMaps sets AttributeMaps field to given value.

### HasAttributeMaps

`func (o *AttributeMapperProfileDTO) HasAttributeMaps() bool`

HasAttributeMaps returns a boolean if a field has been set.

### GetElementId

`func (o *AttributeMapperProfileDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *AttributeMapperProfileDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *AttributeMapperProfileDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *AttributeMapperProfileDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *AttributeMapperProfileDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeMapperProfileDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeMapperProfileDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AttributeMapperProfileDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludeNonMappedProperties

`func (o *AttributeMapperProfileDTO) GetIncludeNonMappedProperties() bool`

GetIncludeNonMappedProperties returns the IncludeNonMappedProperties field if non-nil, zero value otherwise.

### GetIncludeNonMappedPropertiesOk

`func (o *AttributeMapperProfileDTO) GetIncludeNonMappedPropertiesOk() (*bool, bool)`

GetIncludeNonMappedPropertiesOk returns a tuple with the IncludeNonMappedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNonMappedProperties

`func (o *AttributeMapperProfileDTO) SetIncludeNonMappedProperties(v bool)`

SetIncludeNonMappedProperties sets IncludeNonMappedProperties field to given value.

### HasIncludeNonMappedProperties

`func (o *AttributeMapperProfileDTO) HasIncludeNonMappedProperties() bool`

HasIncludeNonMappedProperties returns a boolean if a field has been set.

### GetName

`func (o *AttributeMapperProfileDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeMapperProfileDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeMapperProfileDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttributeMapperProfileDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileType

`func (o *AttributeMapperProfileDTO) GetProfileType() string`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *AttributeMapperProfileDTO) GetProfileTypeOk() (*string, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *AttributeMapperProfileDTO) SetProfileType(v string)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *AttributeMapperProfileDTO) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


