# GetTomcatExecEnvsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**TomcatExecEnv** | Pointer to [**[]TomcatExecutionEnvironmentDTO**](TomcatExecutionEnvironmentDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetTomcatExecEnvsRes

`func NewGetTomcatExecEnvsRes() *GetTomcatExecEnvsRes`

NewGetTomcatExecEnvsRes instantiates a new GetTomcatExecEnvsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTomcatExecEnvsResWithDefaults

`func NewGetTomcatExecEnvsResWithDefaults() *GetTomcatExecEnvsRes`

NewGetTomcatExecEnvsResWithDefaults instantiates a new GetTomcatExecEnvsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetTomcatExecEnvsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetTomcatExecEnvsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetTomcatExecEnvsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetTomcatExecEnvsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTomcatExecEnv

`func (o *GetTomcatExecEnvsRes) GetTomcatExecEnv() []TomcatExecutionEnvironmentDTO`

GetTomcatExecEnv returns the TomcatExecEnv field if non-nil, zero value otherwise.

### GetTomcatExecEnvOk

`func (o *GetTomcatExecEnvsRes) GetTomcatExecEnvOk() (*[]TomcatExecutionEnvironmentDTO, bool)`

GetTomcatExecEnvOk returns a tuple with the TomcatExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTomcatExecEnv

`func (o *GetTomcatExecEnvsRes) SetTomcatExecEnv(v []TomcatExecutionEnvironmentDTO)`

SetTomcatExecEnv sets TomcatExecEnv field to given value.

### HasTomcatExecEnv

`func (o *GetTomcatExecEnvsRes) HasTomcatExecEnv() bool`

HasTomcatExecEnv returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetTomcatExecEnvsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetTomcatExecEnvsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetTomcatExecEnvsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetTomcatExecEnvsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


