# StoreIdSourceDbRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdSourceDb** | Pointer to [**DbIdentitySourceDTO**](DbIdentitySourceDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreIdSourceDbRes

`func NewStoreIdSourceDbRes() *StoreIdSourceDbRes`

NewStoreIdSourceDbRes instantiates a new StoreIdSourceDbRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdSourceDbResWithDefaults

`func NewStoreIdSourceDbResWithDefaults() *StoreIdSourceDbRes`

NewStoreIdSourceDbResWithDefaults instantiates a new StoreIdSourceDbRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreIdSourceDbRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreIdSourceDbRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreIdSourceDbRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreIdSourceDbRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdSourceDb

`func (o *StoreIdSourceDbRes) GetIdSourceDb() DbIdentitySourceDTO`

GetIdSourceDb returns the IdSourceDb field if non-nil, zero value otherwise.

### GetIdSourceDbOk

`func (o *StoreIdSourceDbRes) GetIdSourceDbOk() (*DbIdentitySourceDTO, bool)`

GetIdSourceDbOk returns a tuple with the IdSourceDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSourceDb

`func (o *StoreIdSourceDbRes) SetIdSourceDb(v DbIdentitySourceDTO)`

SetIdSourceDb sets IdSourceDb field to given value.

### HasIdSourceDb

`func (o *StoreIdSourceDbRes) HasIdSourceDb() bool`

HasIdSourceDb returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreIdSourceDbRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreIdSourceDbRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreIdSourceDbRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreIdSourceDbRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


