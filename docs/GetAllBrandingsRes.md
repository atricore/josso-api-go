# GetAllBrandingsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brandings** | Pointer to [**[]BrandingDefinitionDTO**](BrandingDefinitionDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAllBrandingsRes

`func NewGetAllBrandingsRes() *GetAllBrandingsRes`

NewGetAllBrandingsRes instantiates a new GetAllBrandingsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllBrandingsResWithDefaults

`func NewGetAllBrandingsResWithDefaults() *GetAllBrandingsRes`

NewGetAllBrandingsResWithDefaults instantiates a new GetAllBrandingsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandings

`func (o *GetAllBrandingsRes) GetBrandings() []BrandingDefinitionDTO`

GetBrandings returns the Brandings field if non-nil, zero value otherwise.

### GetBrandingsOk

`func (o *GetAllBrandingsRes) GetBrandingsOk() (*[]BrandingDefinitionDTO, bool)`

GetBrandingsOk returns a tuple with the Brandings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandings

`func (o *GetAllBrandingsRes) SetBrandings(v []BrandingDefinitionDTO)`

SetBrandings sets Brandings field to given value.

### HasBrandings

`func (o *GetAllBrandingsRes) HasBrandings() bool`

HasBrandings returns a boolean if a field has been set.

### GetError

`func (o *GetAllBrandingsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAllBrandingsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAllBrandingsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetAllBrandingsRes) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


