# StoreIisExecEnvRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IisExecEnv** | Pointer to [**WindowsIISExecutionEnvironmentDTO**](WindowsIISExecutionEnvironmentDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreIisExecEnvRes

`func NewStoreIisExecEnvRes() *StoreIisExecEnvRes`

NewStoreIisExecEnvRes instantiates a new StoreIisExecEnvRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIisExecEnvResWithDefaults

`func NewStoreIisExecEnvResWithDefaults() *StoreIisExecEnvRes`

NewStoreIisExecEnvResWithDefaults instantiates a new StoreIisExecEnvRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreIisExecEnvRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreIisExecEnvRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreIisExecEnvRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreIisExecEnvRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIisExecEnv

`func (o *StoreIisExecEnvRes) GetIisExecEnv() WindowsIISExecutionEnvironmentDTO`

GetIisExecEnv returns the IisExecEnv field if non-nil, zero value otherwise.

### GetIisExecEnvOk

`func (o *StoreIisExecEnvRes) GetIisExecEnvOk() (*WindowsIISExecutionEnvironmentDTO, bool)`

GetIisExecEnvOk returns a tuple with the IisExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIisExecEnv

`func (o *StoreIisExecEnvRes) SetIisExecEnv(v WindowsIISExecutionEnvironmentDTO)`

SetIisExecEnv sets IisExecEnv field to given value.

### HasIisExecEnv

`func (o *StoreIisExecEnvRes) HasIisExecEnv() bool`

HasIisExecEnv returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreIisExecEnvRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreIisExecEnvRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreIisExecEnvRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreIisExecEnvRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


