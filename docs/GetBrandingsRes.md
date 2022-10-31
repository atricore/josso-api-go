# GetBrandingsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brandings** | Pointer to [**[]BrandingDefinitionDTO**](BrandingDefinitionDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewGetBrandingsRes

`func NewGetBrandingsRes() *GetBrandingsRes`

NewGetBrandingsRes instantiates a new GetBrandingsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBrandingsResWithDefaults

`func NewGetBrandingsResWithDefaults() *GetBrandingsRes`

NewGetBrandingsResWithDefaults instantiates a new GetBrandingsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandings

`func (o *GetBrandingsRes) GetBrandings() []BrandingDefinitionDTO`

GetBrandings returns the Brandings field if non-nil, zero value otherwise.

### GetBrandingsOk

`func (o *GetBrandingsRes) GetBrandingsOk() (*[]BrandingDefinitionDTO, bool)`

GetBrandingsOk returns a tuple with the Brandings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandings

`func (o *GetBrandingsRes) SetBrandings(v []BrandingDefinitionDTO)`

SetBrandings sets Brandings field to given value.

### HasBrandings

`func (o *GetBrandingsRes) HasBrandings() bool`

HasBrandings returns a boolean if a field has been set.

### GetError

`func (o *GetBrandingsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetBrandingsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetBrandingsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetBrandingsRes) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


