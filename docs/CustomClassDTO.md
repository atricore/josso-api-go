# CustomClassDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cardinality** | Pointer to **string** |  | [optional] 
**Fqcn** | Pointer to **string** |  | [optional] 
**OsgiFilter** | Pointer to **string** |  | [optional] 
**OsgiService** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to [**[]CustomClassPropertyDTO**](CustomClassPropertyDTO.md) |  | [optional] 
**TimeoutSecs** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomClassDTO

`func NewCustomClassDTO() *CustomClassDTO`

NewCustomClassDTO instantiates a new CustomClassDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomClassDTOWithDefaults

`func NewCustomClassDTOWithDefaults() *CustomClassDTO`

NewCustomClassDTOWithDefaults instantiates a new CustomClassDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardinality

`func (o *CustomClassDTO) GetCardinality() string`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *CustomClassDTO) GetCardinalityOk() (*string, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *CustomClassDTO) SetCardinality(v string)`

SetCardinality sets Cardinality field to given value.

### HasCardinality

`func (o *CustomClassDTO) HasCardinality() bool`

HasCardinality returns a boolean if a field has been set.

### GetFqcn

`func (o *CustomClassDTO) GetFqcn() string`

GetFqcn returns the Fqcn field if non-nil, zero value otherwise.

### GetFqcnOk

`func (o *CustomClassDTO) GetFqcnOk() (*string, bool)`

GetFqcnOk returns a tuple with the Fqcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqcn

`func (o *CustomClassDTO) SetFqcn(v string)`

SetFqcn sets Fqcn field to given value.

### HasFqcn

`func (o *CustomClassDTO) HasFqcn() bool`

HasFqcn returns a boolean if a field has been set.

### GetOsgiFilter

`func (o *CustomClassDTO) GetOsgiFilter() string`

GetOsgiFilter returns the OsgiFilter field if non-nil, zero value otherwise.

### GetOsgiFilterOk

`func (o *CustomClassDTO) GetOsgiFilterOk() (*string, bool)`

GetOsgiFilterOk returns a tuple with the OsgiFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsgiFilter

`func (o *CustomClassDTO) SetOsgiFilter(v string)`

SetOsgiFilter sets OsgiFilter field to given value.

### HasOsgiFilter

`func (o *CustomClassDTO) HasOsgiFilter() bool`

HasOsgiFilter returns a boolean if a field has been set.

### GetOsgiService

`func (o *CustomClassDTO) GetOsgiService() bool`

GetOsgiService returns the OsgiService field if non-nil, zero value otherwise.

### GetOsgiServiceOk

`func (o *CustomClassDTO) GetOsgiServiceOk() (*bool, bool)`

GetOsgiServiceOk returns a tuple with the OsgiService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsgiService

`func (o *CustomClassDTO) SetOsgiService(v bool)`

SetOsgiService sets OsgiService field to given value.

### HasOsgiService

`func (o *CustomClassDTO) HasOsgiService() bool`

HasOsgiService returns a boolean if a field has been set.

### GetProperties

`func (o *CustomClassDTO) GetProperties() []CustomClassPropertyDTO`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomClassDTO) GetPropertiesOk() (*[]CustomClassPropertyDTO, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomClassDTO) SetProperties(v []CustomClassPropertyDTO)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CustomClassDTO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *CustomClassDTO) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *CustomClassDTO) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *CustomClassDTO) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *CustomClassDTO) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetType

`func (o *CustomClassDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomClassDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomClassDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomClassDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


