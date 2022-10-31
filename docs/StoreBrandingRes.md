# StoreBrandingRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branding** | Pointer to [**CustomBrandingDefinitionDTO**](CustomBrandingDefinitionDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewStoreBrandingRes

`func NewStoreBrandingRes() *StoreBrandingRes`

NewStoreBrandingRes instantiates a new StoreBrandingRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreBrandingResWithDefaults

`func NewStoreBrandingResWithDefaults() *StoreBrandingRes`

NewStoreBrandingResWithDefaults instantiates a new StoreBrandingRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranding

`func (o *StoreBrandingRes) GetBranding() CustomBrandingDefinitionDTO`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *StoreBrandingRes) GetBrandingOk() (*CustomBrandingDefinitionDTO, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *StoreBrandingRes) SetBranding(v CustomBrandingDefinitionDTO)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *StoreBrandingRes) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetError

`func (o *StoreBrandingRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreBrandingRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreBrandingRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreBrandingRes) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


