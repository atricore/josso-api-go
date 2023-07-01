# GetBundlesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundles** | Pointer to [**[]BundleDescr**](BundleDescr.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetBundlesRes

`func NewGetBundlesRes() *GetBundlesRes`

NewGetBundlesRes instantiates a new GetBundlesRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBundlesResWithDefaults

`func NewGetBundlesResWithDefaults() *GetBundlesRes`

NewGetBundlesResWithDefaults instantiates a new GetBundlesRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundles

`func (o *GetBundlesRes) GetBundles() []BundleDescr`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *GetBundlesRes) GetBundlesOk() (*[]BundleDescr, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *GetBundlesRes) SetBundles(v []BundleDescr)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *GetBundlesRes) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetError

`func (o *GetBundlesRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetBundlesRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetBundlesRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetBundlesRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetBundlesRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetBundlesRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetBundlesRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetBundlesRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


